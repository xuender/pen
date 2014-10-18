package class

import (
	"../base"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
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

// 删除老师
func delTeacherEvent(data *string, session base.Session) (ret interface{}, err error) {
	ret = "ok"
	var id int64
	json.Unmarshal([]byte(*data), &id)
	log.WithFields(log.Fields{
		"ID": id,
	}).Debug("删除老师")
	var e ClassTeacher
	err = db.First(&e, id).Error
	if err == nil {
		err = db.Delete(&e).Error
	}
	return
}
func init() {
	base.RegisterEvent(Code, 教师查询, getTeacherEvent)
	base.RegisterEvent(Code, 教师编辑, updateTeacherEvent)
	base.RegisterEvent(Code, 教师删除, delTeacherEvent)
	meta.AddDbFunc(func() {
		db.AutoMigrate(&ClassTeacher{})
	})
}
