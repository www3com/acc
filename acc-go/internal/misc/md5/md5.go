package md5

import (
	"crypto/md5"
	"fmt"
)

func Encrypt(value string) string {
	data := []byte(value) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func Valid(value string, encryptedValue string) bool {
	encrypt := Encrypt(value)
	if encrypt == encryptedValue {
		return true
	}
	return false
}
