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
