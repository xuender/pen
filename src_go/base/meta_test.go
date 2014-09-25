package base

import (
	"testing"
)

//获取元数据
func TestGetMeta(t *testing.T) {
	_, err := GetMeta("error")
	if err == nil {
		t.Errorf("error错误")
	} else {
		meta, err := GetMeta("base")
		if err != nil {
			t.Errorf("base没有找到")
		}
		if meta.Code != "base" {
			t.Errorf("code错误")
		}
	}
}

var me = Meta{"测试功能", "test", "测试时使用的元数据", []uint{1, 5, 7}}

// 判断权限
func TestHasAction(t *testing.T) {
	if me.hasAction(2, 5) {
		t.Errorf("没有权限5")
	}
	if !me.hasAction(34, 5) {
		t.Errorf("有权限5")
	}
	if !me.hasAction(32, 5) {
		t.Errorf("有权限5")
	}
	if me.hasAction(32, 7) {
		t.Errorf("没有权限7")
	}
}

// 所有权限
func TestAllAction(t *testing.T) {
	if me.allAction() != 2+32+128 {
		t.Errorf("allAction方法错误")
	}
}

// 增加权限
func TestAddAction(t *testing.T) {
	if me.addAction(2, 5, 7) != 2+32+128 {
		t.Errorf("addAction方法错误")
	}
	if me.addAction(34, 7) != 2+32+128 {
		t.Errorf("addAction方法错误")
	}
	if me.addAction(34, 7, 9) != 2+32+128 {
		t.Errorf("addAction增加不存在方法错误")
	}
}
