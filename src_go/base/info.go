package base

// 功能定义
//import (
//	log "github.com/Sirupsen/logrus"
//)

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

var meta = Meta{Name: "基本功能", Code: Code, Description: "用户管理、身份认证", Action: map[uint]string{
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
}}

func init() {
	RegisterMeta(&meta)
}
