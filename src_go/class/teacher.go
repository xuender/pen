package class

import (
	"../base"
	"encoding/json"
	"time"
)

// 教师
type ClassTeacher struct {
	base.BaseObject
	ClassPeople `sql:"-"`
	PeopleId    int64
	// 执教日期
	StartAt time.Time
}

// 获取教师信息列表
func getTeacherEvent(data *string, session base.Session) (interface{}, error) {
	var vs []ClassTeacher
	db.Find(&vs)
	var ret []ClassTeacher
	for _, v := range vs {
		db.First(&v.ClassPeople, v.PeopleId)
		ret = append(ret, v)
	}
	return ret, nil
}
func updateTeacherEvent(data *string, session base.Session) (ret interface{}, err error) {
	//TODO 权限认证
	ret = "ok"
	var e ClassTeacher
	json.Unmarshal([]byte(*data), &e)
	if e.Id == 0 { //增加
		p := e.ClassPeople
		db.Save(&p)
		e.PeopleId = p.Id
		err = db.Save(&e).Error
	} else { //修改
		var o ClassTeacher
		var p ClassPeople
		db.First(&o, e.Id)
		db.Save(&e)
		db.First(&p, e.PeopleId)
		p = e.ClassPeople
		p.Id = e.PeopleId
		err = db.Save(&p).Error
	}
	return
}
func init() {
	base.RegisterEvent(Code, 教师, getTeacherEvent)
	base.RegisterEvent(Code, 编辑教师, updateTeacherEvent)
	db.AutoMigrate(&ClassTeacher{})
}
