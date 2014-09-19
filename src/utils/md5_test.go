package utils

import "testing"

func TestMd5(t *testing.T) {
	if Md5("123") != "202cb962ac59075b964b07152d234b70" {
		t.Errorf("字符串DM5失败:%s", Md5("123"))
	}
	v, _ := Md5File("/dev/null")
	if v != "d41d8cd98f00b204e9800998ecf8427e" {
		t.Errorf("文件DM5失败:%s", v)
	}
}
