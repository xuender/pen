package class

import (
	"../base"
	//"database/sql"
	"code.google.com/p/go.net/websocket"
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
func getEmployeeEvent(data *string, ws *websocket.Conn, session base.Session) {
	var vs []ClassEmployee
	db.Find(&vs)
	var ret []ClassEmployee
	for _, v := range vs {
		v.read()
		ret = append(ret, v)
	}
	base.Send(ws, Code, 雇员, ret)
}
func updateEmployeeEvent(data *string, ws *websocket.Conn, session base.Session) {
	//TODO 权限认证
	var e ClassEmployee
	json.Unmarshal([]byte(*data), &e)
	log.WithFields(log.Fields{
		"ID": e.Id,
	}).Debug("修改雇员")
	if e.Id == 0 { //增加
		p := e.ClassPeople
		db.Save(&p)
		e.PeopleId = p.Id
		err := db.Save(&e).Error
		if err == nil {
			base.Send(ws, Code, 编辑雇员, "ok")
		} else {
			log.Debug(err)
			base.Send(ws, base.Code, base.MSG, err)
		}
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
		err := db.Save(&p).Error
		if err == nil {
			base.Send(ws, Code, 编辑雇员, "ok")
		} else {
			log.Debug(err)
			base.Send(ws, base.Code, base.MSG, err)
		}
	}
}
func init() {
	base.RegisterEvent(Code, 雇员, getEmployeeEvent)
	base.RegisterEvent(Code, 编辑雇员, updateEmployeeEvent)
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
