package base

// 功能定义
import (
	log "github.com/Sirupsen/logrus"
)

const (
	消息 = iota
	登录
	登出
	人数
	用户列表
	字典
	字典版本
	查看字典
	修改字典
	Code = "base"
)

func init() {
	log.WithFields(log.Fields{
		"消息":   消息,
		"登录":   登录,
		"登出":   登出,
		"人数":   人数,
		"用户列表": 用户列表,
		"字典":   字典,
		"字典版本": 字典版本,
		"查看字典": 查看字典,
		"修改字典": 修改字典,
	}).Debug("枚举")
	RegisterMeta(Meta{"基本功能", "base", "用户管理、身份认证", []uint{
		登录,
		登出,
		人数,
		用户列表,
		字典,
		字典版本,
	}})
}
