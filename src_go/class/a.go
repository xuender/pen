package class

import (
	"../base"
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

const (
	班级 = iota
	雇员
	编辑雇员
	教师
	编辑教师
	学员
	Code = "class"
)

var db *gorm.DB

func init() {
	log.WithFields(log.Fields{
		"雇员":   雇员,
		"编辑雇员": 编辑雇员,
		"教师":   教师,
		"编辑教师": 编辑教师,
		"班级":   班级,
		"学员":   学员,
	}).Debug("class 枚举")
	base.RegisterMeta(base.Meta{"学习班", Code, "学习班管理", map[uint]string{
		雇员:   "雇员",
		编辑雇员: "编辑雇员",
		班级:   "班级",
		教师:   "教师",
		编辑教师: "编辑教师",
		学员:   "学员",
	}})
	db = base.Db()
}
