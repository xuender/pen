package class

import (
	"../base"
	"code.google.com/p/go.net/websocket"
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
func getTeacherEvent(data *string, ws *websocket.Conn, session base.Session) {
	var vs []ClassTeacher
	db.Find(&vs)
	var ret []ClassTeacher
	for _, v := range vs {
		db.First(&v.ClassPeople, v.PeopleId)
		ret = append(ret, v)
	}
	base.Send(ws, Code, 教师, ret)
}
func updateTeacherEvent(data *string, ws *websocket.Conn, session base.Session) {
	//TODO 权限认证
	var e ClassTeacher
	json.Unmarshal([]byte(*data), &e)
	if e.Id == 0 { //增加
		p := e.ClassPeople
		db.Save(&p)
		e.PeopleId = p.Id
		err := db.Save(&e).Error
		if err == nil {
			base.Send(ws, Code, 编辑教师, "ok")
		} else {
			log.Debug(err)
			base.Send(ws, base.Code, base.MSG, err)
		}
	} else { //修改
		var o ClassTeacher
		var p ClassPeople
		db.First(&o, e.Id)
		db.Save(&e)
		db.First(&p, e.PeopleId)
		p = e.ClassPeople
		p.Id = e.PeopleId
		err := db.Save(&p).Error
		if err == nil {
			base.Send(ws, Code, 编辑教师, "ok")
		} else {
			log.Debug(err)
			base.Send(ws, base.Code, base.MSG, err)
		}
	}
}
func init() {
	base.RegisterEvent(Code, 教师, getTeacherEvent)
	base.RegisterEvent(Code, 编辑教师, updateTeacherEvent)
	db.AutoMigrate(&ClassTeacher{})
}
