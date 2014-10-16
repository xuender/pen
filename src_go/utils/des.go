package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"strconv"
)

// 加密
func Encrypt(orig, password string) (string, error) {
	in := []byte(orig)
	key := []byte(password[0:8])
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
func Decrypt(crypte, password string) (string, error) {
	in := make([]byte, len(crypte)/2)
	for i := 0; i < len(crypte); i += 2 {
		n, _ := strconv.ParseInt(crypte[i:i+2], 16, 64)
		in[i/2] = byte(n)
	}
	key := []byte(password[0:8])
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCBCDecrypter(block, key)
	out := make([]byte, len(in))
	decrypter.CryptBlocks(out, in)
	out = unpadding(out)
	return string(out), nil
}
func unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
