package class

import (
	"../base"
	//log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

const (
	班级查询 = iota
	班级编辑
	班级删除
	雇员查询
	雇员编辑
	雇员删除
	教师查询
	教师编辑
	教师删除
	学员查询
	学员编辑
	学员删除
	Code = "class"
)

var db *gorm.DB
var meta = base.Meta{Name: "学习班", Code: Code, Description: "学习班管理", Action: map[uint]string{
	班级查询: "班级查询",
	班级编辑: "班级编辑",
	班级删除: "班级删除",
	雇员查询: "雇员查询",
	雇员编辑: "雇员编辑",
	雇员删除: "雇员删除",
	教师查询: "教师查询",
	教师编辑: "教师编辑",
	教师删除: "教师删除",
	学员查询: "学员查询",
	学员编辑: "学员编辑",
	学员删除: "学员删除",
}}

func init() {
	base.RegisterMeta(meta)
	db = base.Db()
}
