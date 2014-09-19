package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func Md5(str string) string {
	// 字符串摘要
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func Md5File(file string) (string, error) {
	// 文件摘要
	f, err := os.Open(file) //打开文件
	defer f.Close()         //打开文件出错处理
	if err != nil {
		return "", err
	}
	h := md5.New()
	buff := bufio.NewReader(f) //读入缓存
	for {
		line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		h.Write([]byte(line))
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
