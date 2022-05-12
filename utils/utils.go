package utils

import (
	"fmt"
	"math/rand"
	"time"
)

type GottaUtils byte

// CreatePwd 生成随机密码,也可以用来生成随机字符串
func (gu *GottaUtils) CreatePwd(length int) string {
	if length <= 0 {
		fmt.Println("密码长度不能为0或负数")
		return ""
	}
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
