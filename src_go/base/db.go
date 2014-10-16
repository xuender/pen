package base

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

// 基础对象
type BaseObject struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//func (u *BaseObject) BeforeCreate() (err error) {
//	u.CreatedAt = time.Now()
//	u.UpdatedAt = u.CreatedAt
//	return
//}
//
//func (u *BaseObject) BeforeUpdate() (err error) {
//	u.UpdatedAt = time.Now()
//	return
//}

// 全局数据库操作对象
var db gorm.DB

// 获取数据库
func Db() *gorm.DB {
	return &db
}

// 初始化
func init() {
	log.WithFields(log.Fields{
		"dialect": PenConfig.Db.Dialect,
		"source":  PenConfig.Db.GetSource(),
	}).Debug("db open")
	d, err := gorm.Open(PenConfig.Db.Dialect, PenConfig.Db.GetSource())
	if err != nil {
		panic(err)
	}
	d.LogMode(true)
	d.DB().SetMaxIdleConns(10)
	d.DB().SetMaxOpenConns(100)
	d.SingularTable(true)
	db = d
}
