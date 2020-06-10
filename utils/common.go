package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

//返回一个长度为n的随机字符串
func RandomStringRune(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	c := make([]rune, n)
	for i := range c {
		c[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(c)
}

//创建token
func GenToken(telephone string) string {
	ts := fmt.Sprintf("%x", time.Now().Unix())
	m := md5.New()
	_, _ = m.Write([]byte(telephone + ts))
	tokenPrefix := m.Sum(nil)
	//128+8
	return hex.EncodeToString(tokenPrefix) + ":" + ts[:8]

}
