package base

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	//	"fmt"
	log "github.com/Sirupsen/logrus"
)

// 字典数据
type Dict struct {
	BaseObject
	// 类型
	Type string
	// 代码
	Code string
	// 标题
	Title string
}

// 字典版本
type DictVer struct {
	BaseObject
	// 类型
	Type string
	// 版本
	Ver int
}

// 字典信息
type DictMessage struct {
	Type string            `json:"type"`
	Ver  int               `json:"ver"`
	Dict map[string]string `json:"dict"`
}

// 查询字典信息
func (dm *DictMessage) read() {
	var ds []Dict
	db.Where("type = ?", dm.Type).Find(&ds)
	m := make(map[string]string)
	for _, d := range ds {
		log.Debug(d.Title)
		m[d.Code] = d.Title
	}
	var dv DictVer
	db.Where("type = ?", dm.Type).First(&dv)
	dm.Dict = m
	dm.Ver = dv.Ver
}

// 保存字典，修改字典版本
func (u *Dict) BeforeSave() (err error) {
	var dv DictVer
	db.Where("type = ?", u.Type).First(&dv)
	if dv.Id == 0 {
		dv.Type = u.Type
		dv.Ver = 1
	} else {
		dv.Ver++
	}
	dictMap[u.Type] = dv.Ver
	db.Save(&dv)
	return
}

// 字典保存后自动推送所有客户端
func (u *Dict) publish() {
	var tm DictMessage
	tm.Type = u.Type
	tm.read()
	//tm.Dict[u.Code] = u.Title
	for ws, session := range onlines {
		if session.IsLogin {
			send(ws, Code, 字典, tm)
		}
	}
}

var dictMap = make(map[string]int)

//字典版本对比事件
func dictVerEvent(data *string, ws *websocket.Conn, session Session) {
	m := make(map[string]int)
	json.Unmarshal([]byte(*data), &m)
	for k, v := range dictMap {
		u, ok := m[k]
		if ok && u == v {
			continue
		}
		go dictSend(k, ws)
	}
}
func dictSend(t string, ws *websocket.Conn) {
	log.WithFields(log.Fields{
		"type": t,
	}).Debug("发送字典")
	var tm DictMessage
	tm.Type = t
	tm.read()
	send(ws, Code, 字典, tm)
}

// 查看字典
func getDictEvent(data *string, ws *websocket.Conn, session Session) {
	var ds []Dict
	db.Where("type = ?", data).Find(&ds)
	send(ws, Code, 查看字典, ds)
}

// 修改字典
func updateDictEvent(data *string, ws *websocket.Conn, session Session) {
	var d Dict
	json.Unmarshal([]byte(*data), &d)
	log.WithFields(log.Fields{
		"ID": d.Id,
	}).Debug("update Dict Event")
	if d.Id == 0 {
		e := db.Save(&d).Error
		if e == nil {
			d.publish()
		} else {
			log.Debug(e)
			send(ws, Code, 消息, e)
		}
	} else {
		var o Dict
		db.First(&o, d.Id)
		o.Code = d.Code
		o.Title = d.Title
		e := db.Save(&o).Error
		if e == nil {
			o.publish()
		} else {
			log.Debug(e)
			send(ws, Code, 消息, e)
		}
	}
}

// 初始化
func init() {
	RegisterEvent(Code, 字典版本, dictVerEvent)
	RegisterEvent(Code, 查看字典, getDictEvent)
	RegisterEvent(Code, 修改字典, updateDictEvent)
	db.AutoMigrate(&Dict{}, &DictVer{})
	db.Model(&Dict{}).AddUniqueIndex("idx_dict_code", "type", "code")
	errorMessage["idx_dict_code"] = "字典代码不能重复"
	db.Model(&DictVer{}).AddUniqueIndex("idx_dict_ver", "type")
	var dvs []DictVer
	db.Find(&dvs)
	for _, dv := range dvs {
		dictMap[dv.Type] = dv.Ver
	}
	//db.Create(&Dict{
	//	Type:  "type",
	//	Code:  "gender",
	//	Title: "性别",
	//})
	//db.Create(&Dict{
	//	Type:  "type",
	//	Code:  "province",
	//	Title: "省份",
	//})
	//db.Create(&Dict{
	//	Type:  "gender",
	//	Code:  "M",
	//	Title: "男",
	//})
	//db.Create(&Dict{
	//	Type:  "gender",
	//	Code:  "F",
	//	Title: "女",
	//})
	//db.Create(&Dict{
	//	Type:  "province",
	//	Code:  "SD",
	//	Title: "山东省",
	//})
}
