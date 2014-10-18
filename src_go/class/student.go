package class

import (
	"../base"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"time"
)

// 学生
type ClassStudent struct {
	base.BaseObject
	ClassPeople `sql:"-"`
	PeopleId    int64
	// 入学日期
	StartAt time.Time
}

// 获取学生信息列表
func getStudentEvent(data *string, session base.Session) (interface{}, error) {
	var vs []ClassStudent
	db.Find(&vs)
	var ret []ClassStudent
	for _, v := range vs {
		db.First(&v.ClassPeople, v.PeopleId)
		ret = append(ret, v)
	}
	return ret, nil
}
func updateStudentEvent(data *string, session base.Session) (ret interface{}, err error) {
	ret = "ok"
	var e ClassStudent
	json.Unmarshal([]byte(*data), &e)
	if e.Id == 0 { //增加
		p := e.ClassPeople
		db.Save(&p)
		e.PeopleId = p.Id
		err = db.Save(&e).Error
	} else { //修改
		var o ClassStudent
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

// 删除学员
func delStudentEvent(data *string, session base.Session) (ret interface{}, err error) {
	ret = "ok"
	var id int64
	json.Unmarshal([]byte(*data), &id)
	log.WithFields(log.Fields{
		"ID": id,
	}).Debug("删除学员")
	var e ClassStudent
	err = db.First(&e, id).Error
	if err == nil {
		err = db.Delete(&e).Error
	}
	return
}
func init() {
	base.RegisterEvent(Code, 学员查询, getStudentEvent)
	base.RegisterEvent(Code, 学员编辑, updateStudentEvent)
	base.RegisterEvent(Code, 学员删除, delStudentEvent)
	meta.AddDbFunc(func() {
		db.AutoMigrate(&ClassStudent{})
	})
}
