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
	Dict map[string]string `json:"data"`
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
	var ds []Dict
	db.Where("type = ?", t).Find(&ds)
	m := make(map[string]string)
	for _, d := range ds {
		m[d.Code] = d.Title
	}
	//d, err := json.Marshal(m)
	//if err != nil {
	//  log.WithFields(log.Fields{
	//    "JSON": m,
	//    "err":  err,
	//  }).Error("JSON编码错误")
	//  return
	//}
	var dv DictVer
	db.Where("type = ?", t).First(&dv)
	var tm DictMessage
	tm.Type = t
	tm.Dict = m
	tm.Ver = dv.Ver
	send(ws, Code, 字典, tm)
}

// 初始化
func init() {
	RegisterEvent(Code, 字典版本, dictVerEvent)
	db.AutoMigrate(&Dict{}, &DictVer{})
	db.Model(&Dict{}).AddUniqueIndex("idx_dict_code", "type", "code")
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
