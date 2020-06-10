package utils

import (
	"math/rand"
	"time"
)

//返回一个长度为n的随机字符串
func GenSalt(n int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	c := make([]rune, n)
	for i := range c {
		c[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(c)
}
