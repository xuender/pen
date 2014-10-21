package base

import (
	"errors"
	"fmt"
	gd "github.com/xuender/godecode2"
	"strings"
	//  log "github.com/Sirupsen/logrus"
)

// 模块元信息
type Meta struct {
	Name          string
	Code          string
	Description   string
	Action        map[uint]string
	DbCreateFuncs []func()
	DbInitFuncs   []func()
}

// 增加数据库初始化方法
func (m *Meta) AddDbFunc(f func()) {
	m.DbCreateFuncs = append(m.DbCreateFuncs, f)
}

// 增加数据库启动执行的方法
func (m *Meta) AddDbInitFunc(f func()) {
	m.DbInitFuncs = append(m.DbInitFuncs, f)
}

var metas []*Meta

// 注册元信息
func RegisterMeta(meta *Meta) {
	metas = append(metas, meta)
}

// 获取注册元信息
func GetMetaJs() string {
	ret := ""
	for _, m := range metas {
		ret += "var " + strings.ToUpper(m.Code) + "={"
		for k, v := range m.Action {
			ret += fmt.Sprintf("%s:%d,", gd.Initials(v), k)
		}
		ret += "};\n"
	}
	return ret
}

// 全部功能取值
func (r Meta) allAction() int64 {
	var ret int64 = 0
	for i, _ := range r.Action {
		ret = ret | int64(1<<i)
	}
	return ret
}

// 判断是否具有功能
func (r Meta) hasAction(power int64, actions ...uint) bool {
	// 判断actions是否包含在Meta中
	for _, a := range actions {
		no := true
		for i, _ := range r.Action {
			if a == i {
				no = false
			}
		}
		if no {
			return false
		}
	}
	// 判断是否包含功能
	var tmp int64 = 0
	for _, a := range actions {
		tmp = tmp | int64(1<<a)
	}
	return power == power|tmp
}

// 增加功能权限
func (r Meta) addAction(power int64, actions ...uint) int64 {
	for _, a := range actions {
		for i, _ := range r.Action {
			if a == i {
				power = power | int64(1<<a)
			}
		}
	}
	return power
}

// 查找元信息
func GetMeta(code string) (*Meta, error) {
	for _, m := range metas {
		if m.Code == code {
			return m, nil
		}
	}
	return nil, errors.New("code没有找到")
}
