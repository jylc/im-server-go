package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"im-server-go/connections"
	"im-server-go/entity/mysql_model"
	"im-server-go/utils"
	"log"
	"strings"
)

//登录
func Login(telephone, password string) (mysql_model.User, error) {
	conn := connections.GetSqlConn()
	var user mysql_model.User
	var err error
	if conn == nil {
		log.Println("[Service Login] Get Sql Connection error")
		return user, errors.New("")
	}

	user, err = conn.GetUserInfo(telephone)
	if err != nil {
		log.Println("[Service Login] ", err)
		return mysql_model.User{}, err
	}
	p := user.Password
	splitStorage := strings.Split(p, ":")
	hash := sha1.New()
	_, err = hash.Write([]byte(password + splitStorage[0]))
	bs := hex.EncodeToString(hash.Sum(nil))
	if err != nil {
		return user, err
	}
	bs = splitStorage[0] + ":" + bs //加密后的密码

	//用户登录时创建token
	token := utils.GenToken(telephone)
	if conn.UserSignIn(telephone, bs, token) == true {
		return user, nil
	} else {
		return user, errors.New("login failed")
	}
}
