package class

import (
	"../base"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"time"
)

// 班级
type Class struct {
	base.BaseObject
	// 班级名称
	Name string
	// 起始日期
	StartAt time.Time
	// 结束日期
	EndAt time.Time
}

// 获取义工信息列表
func getClassEvent(data *string, session base.Session) (interface{}, error) {
	var cs []Class
	db.Find(&cs)
	return cs, nil
}

// 班级修改
func updateClassEvent(data *string, session base.Session) (ret interface{}, err error) {
	ret = "ok"
	var e Class
	json.Unmarshal([]byte(*data), &e)
	log.WithFields(log.Fields{
		"ID": e.Id,
	}).Debug("班级修改")
	err = db.Save(&e).Error
	return
}

// 班级删除
func delClassEvent(data *string, session base.Session) (ret interface{}, err error) {
	ret = "ok"
	var id int64
	json.Unmarshal([]byte(*data), &id)
	log.WithFields(log.Fields{
		"ID": id,
	}).Debug("删除雇员")
	var e Class
	err = db.First(&e, id).Error
	if err == nil {
		err = db.Delete(&e).Error
	}
	return
}
func init() {
	log.WithFields(log.Fields{
		"Name": "Class",
	}).Info("init")
	base.RegisterEvent(Code, 班级查询, getClassEvent)
	base.RegisterEvent(Code, 班级编辑, updateClassEvent)
	base.RegisterEvent(Code, 班级删除, delClassEvent)
	meta.AddDbFunc(func() {
		log.WithFields(log.Fields{
			"Name": "Class",
		}).Debug("表格")
		if db.CreateTable(&Class{}).Error == nil {
			db.Save(&Class{Name: "三十一期1班"})
			db.Save(&Class{Name: "三十一期2班"})
		} else {
			db.AutoMigrate(&Class{})
		}
	})
}
