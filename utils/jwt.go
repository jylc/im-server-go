package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

//创建token
func GenToken(telephone string) string {
	ts := fmt.Sprintf("%x", time.Now().Unix())
	m := md5.New()
	_, _ = m.Write([]byte(telephone + ts))
	tokenPrefix := m.Sum(nil)
	//128+8
	return hex.EncodeToString(tokenPrefix) + ":" + ts[:8]
}
