package base
import (
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "time"
)

// 基础对象
type BaseObject struct {
  Id        int64
  Created   time.Time
  Updated   time.Time
}

func (u *BaseObject) BeforeCreate() (err error) {
  u.Created = time.Now()
  u.Updated = u.Created
  return
}

func (u *BaseObject) BeforeUpdate() (err error) {
  u.Updated = time.Now()
  return
}

var db gorm.DB
func init(){
  d, err := gorm.Open("postgres", "user=postgres dbname=go password=xcy123 sslmode=disable")
  if err != nil {
    panic(err)
  }
  d.LogMode(true)
  d.DB().SetMaxIdleConns(10)
  d.DB().SetMaxOpenConns(100)
  d.SingularTable(true)
  db=d
}
