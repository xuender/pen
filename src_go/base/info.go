package base

// 功能定义
import (
	log "github.com/Sirupsen/logrus"
)

const (
	MSG = iota
	登录
	登出
	人数
	用户查询
	用户编辑
	字典获取
	字典版本
	字典查询
	字典编辑
	Code = "base"
)

func init() {
	log.WithFields(log.Fields{
		"MSG":  MSG,
		"登录":   登录,
		"登出":   登出,
		"人数":   人数,
		"用户查询": 用户查询,
		"用户编辑": 用户编辑,
		"字典获取": 字典获取,
		"字典版本": 字典版本,
		"字典查询": 字典查询,
		"字典编辑": 字典编辑,
	}).Debug("枚举")
	RegisterMeta(Meta{"基本功能", Code, "用户管理、身份认证", map[uint]string{
		MSG:  "MSG",
		登录:   "登录",
		登出:   "登出",
		人数:   "人数",
		用户查询: "用户查询",
		用户编辑: "用户编辑",
		字典获取: "字典获取",
		字典版本: "字典版本",
		字典查询: "字典查询",
		字典编辑: "字典编辑",
	}})
}
