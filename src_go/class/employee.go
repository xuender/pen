package class

import (
	"../base"
	//"database/sql"
	//"code.google.com/p/go.net/websocket"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
)

// 雇员
type ClassEmployee struct {
	base.BaseObject
	ClassPeople `sql:"-"`
	PeopleId    int64 // Foreign key of ClassPeople
	// 职责
	Duty string
}

// 读取关联的用户信息
func (v *ClassEmployee) read() {
	if v.PeopleId == 0 {
		return
	}
	db.First(&v.ClassPeople, v.PeopleId)
}

// 获取义工信息列表
func getEmployeeEvent(data *string, session base.Session) (interface{}, error) {
	var vs []ClassEmployee
	db.Find(&vs)
	var ret []ClassEmployee
	for _, v := range vs {
		v.read()
		ret = append(ret, v)
	}
	return ret, nil
}
func updateEmployeeEvent(data *string, session base.Session) (ret interface{}, err error) {
	ret = "ok"
	var e ClassEmployee
	json.Unmarshal([]byte(*data), &e)
	log.WithFields(log.Fields{
		"ID": e.Id,
	}).Debug("修改雇员")
	if e.Id == 0 { //增加
		p := e.ClassPeople
		db.Save(&p)
		e.PeopleId = p.Id
		err = db.Save(&e).Error
	} else { //修改
		var o ClassEmployee
		var p ClassPeople
		db.First(&o, e.Id)
		db.Save(&e)
		db.First(&p, o.PeopleId)
		p = e.ClassPeople
		p.Id = e.PeopleId
		log.WithFields(log.Fields{
			"ID":   p.Id,
			"Name": p.Name,
			"City": p.City,
		}).Debug("编辑人员信息")
		err = db.Save(&p).Error
	}
	return
}

// 删除雇员
func delEmployeeEvent(data *string, session base.Session) (ret interface{}, err error) {
	ret = "ok"
	var id int64
	json.Unmarshal([]byte(*data), &id)
	log.WithFields(log.Fields{
		"ID": id,
	}).Debug("删除雇员")
	var e ClassEmployee
	err = db.First(&e, id).Error
	if err == nil {
		err = db.Delete(&e).Error
	}
	return
}
func init() {
	base.RegisterEvent(Code, 雇员查询, getEmployeeEvent)
	base.RegisterEvent(Code, 雇员编辑, updateEmployeeEvent)
	base.RegisterEvent(Code, 雇员删除, delEmployeeEvent)
	if db.CreateTable(&ClassEmployee{}).Error == nil {
		log.WithFields(log.Fields{
			"name": "ClassVolunteer",
		}).Debug("初始化表")
		db.Save(&base.Dict{Type: "type", Code: "duty", Title: "职务"})
		db.Save(&base.Dict{Type: "duty", Code: "CS", Title: "厨师"})
		db.Save(&base.Dict{Type: "duty", Code: "BJ", Title: "保洁"})
		db.Save(&base.Dict{Type: "duty", Code: "ZW", Title: "杂务"})
	} else {
		db.AutoMigrate(&ClassEmployee{})
	}
	//db.Save(&ClassPeople{Name:"测试"})
	//db.Save(&ClassPeople{Name:"王五"})
	//db.Save(&ClassEmployee{Duty:"BJ"})
	//db.Save(&ClassEmployee{Duty:"CS"})
}
