package base
import (
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)


var db gorm.DB
var dbInit = true
// 获取数据库对象
func GetDb() gorm.DB{
  if dbInit {
    dbInit = false
    d, err := gorm.Open("postgres", "user=postgres dbname=go password=xcy123 sslmode=disable")
    if err != nil {
      panic(err)
    }
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
    d.SingularTable(true)
    db = d
  }
  return db
}
