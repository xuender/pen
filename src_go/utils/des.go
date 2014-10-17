package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"fmt"
	"strconv"
)

func getPassword(password string) []byte {
	if len(password) > 7 {
		return []byte(password[0:8])
	}
	if len(password) == 0 {
		return []byte("xuender@")
	}
	o := make([]byte, 8)
	f := 0
	p := []byte(password)
	for i := 0; i < 8; i++ {
		o[i] = p[f]
		f++
		if f >= len(p) {
			f = 0
		}
	}
	return o
}

// 加密
func Encrypt(orig, password string) (string, error) {
	in := []byte(orig)
	key := getPassword(password)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	in = blockAdding(in, block.BlockSize())
	encrypter := cipher.NewCBCEncrypter(block, key)
	out := make([]byte, len(in))
	encrypter.CryptBlocks(out, in)
	return fmt.Sprintf("%X", out), nil
}
func blockAdding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 解密
func Decrypt(crypted, password string) (ret string, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = errors.New("无法解密")
		}
	}()
	in := make([]byte, len(crypted)/2)
	for i := 0; i < len(crypted); i += 2 {
		n, _ := strconv.ParseInt(crypted[i:i+2], 16, 64)
		in[i/2] = byte(n)
	}
	key := getPassword(password)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCBCDecrypter(block, key)
	out := make([]byte, len(in))
	decrypter.CryptBlocks(out, in)
	out = unpadding(out)
	ret = string(out)
	return
}
func unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
