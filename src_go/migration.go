package main
import (
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "./base"
)

func main(){
  db, err := gorm.Open("postgres", "user=postgres dbname=go password=xcy123 sslmode=disable")
  if err != nil {
    panic(err)
  }
  db.SingularTable(true)
  db.AutoMigrate(&base.User{})
  db.Model(&base.User{}).AddUniqueIndex("idx_user_nick", "nick")
}
