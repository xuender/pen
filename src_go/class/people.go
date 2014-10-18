package class

import (
	"../base"
)

type ClassPeople struct {
	base.BaseObject
	Name     string
	Gender   string `sql:type:char(1)`
	Country  string `sql:type:char(3)`
	Province string `sql:type:varchar(10)`
	City     string `sql:type:varchar(40)`
}

func init() {
	meta.AddDbFunc(func() {
		if db.CreateTable(&ClassPeople{}).Error == nil {
			db.Save(&base.Dict{Type: "type", Code: "province", Title: "省份"})
			db.Save(&base.Dict{Type: "province", Code: "SD", Title: "山东省"})
			db.Save(&base.Dict{Type: "province", Code: "BJ", Title: "北京市"})

			db.Save(&base.Dict{Type: "type", Code: "country", Title: "国家地区"})
			db.Save(&base.Dict{Type: "country", Code: "CHN", Title: "中国"})
			db.Save(&base.Dict{Type: "country", Code: "TWN", Title: "中国台湾"})
			db.Save(&base.Dict{Type: "country", Code: "HKG", Title: "中国香港"})
			db.Save(&base.Dict{Type: "country", Code: "USA", Title: "美国"})
			db.Save(&base.Dict{Type: "country", Code: "JPN", Title: "日本"})
			db.Save(&base.Dict{Type: "country", Code: "KOR", Title: "韩国"})
			db.Save(&base.Dict{Type: "country", Code: "MAL", Title: "马来西亚"})
			db.Save(&base.Dict{Type: "country", Code: "OTH", Title: "其他"})
		} else {
			db.AutoMigrate(&ClassPeople{})
		}
	})
}
