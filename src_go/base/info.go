package base

// 功能定义
import (
	log "github.com/Sirupsen/logrus"
)

const (
	登录 = iota
	登出
	人数
	用户列表
	Code = "base"
)

func init() {
	log.WithFields(log.Fields{
		"登录":   登录,
		"登出":   登出,
		"人数":   人数,
		"用户列表": 用户列表,
	}).Debug("枚举")
	RegisterMeta(Meta{"基本功能", "base", "用户管理、身份认证", []uint{
		登录,
		登出,
		人数,
	}})
}
