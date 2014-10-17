package utils

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	s, e := Encrypt("123", "123321321")
	if e != nil {
		t.Errorf("加密异常")
	}
	if s != "B2A8EA7C41E2DC65" {
		t.Errorf("加密错误")
	}
}
func TestDecrypt(t *testing.T) {
	s, e := Decrypt("B2A8EA7C41E2DC65", "123321321")
	if e != nil {
		t.Errorf("解密异常")
	}
	if s != "123" {
		t.Errorf("解密错误")
	}
}
func TestGetPassword(t *testing.T) {
	p := getPassword("1")
	if string(p) != "11111111" {
		t.Errorf("getPassword错误")
	}
	p = getPassword("12")
	if string(p) != "12121212" {
		t.Errorf("getPassword错误")
	}
	p = getPassword("")
	if string(p) != "xuender@" {
		t.Errorf("getPassword错误")
	}
}
func TestError(t *testing.T) {
	_, e := Decrypt("165", "123321321")
	if e == nil {
		t.Errorf("解密异常")
	}
	_, e = Decrypt("65", "123321321")
	if e == nil {
		t.Errorf("解密异常")
	}
}
