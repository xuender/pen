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

// 全局数据库操作对象
var db gorm.DB

// 获取数据库
func Db() *gorm.DB {
	return &db
}
func InitDb() {
	log.WithFields(log.Fields{
		"dialect": BaseConfig.Db.Dialect,
	}).Debug("InitDb")
	d, err := gorm.Open(BaseConfig.Db.Dialect, BaseConfig.Db.GetSource())
	if err != nil {
		panic(err)
	}
	d.LogMode(true)
	d.DB().SetMaxIdleConns(10)
	d.DB().SetMaxOpenConns(100)
	d.SingularTable(true)
	db = d
	for _, m := range metas {
		log.WithFields(log.Fields{
			"code":   m.Code,
			"Action": m.Action,
		}).Debug("meta")
	}
	if BaseConfig.Db.Init {
		BaseConfig.Db.Init = false
		BaseConfig.Save()
		for _, m := range metas {
			for _, f := range m.DbFuncs {
				log.Debug("func")
				f()
			}
		}
	}
}
