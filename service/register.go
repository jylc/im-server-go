package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	dbConn "im-server-go/connections"
	"im-server-go/entity/mysql_model"
	"im-server-go/utils"
	"log"
)

func Register(telephone, pwd string) (mysql_model.User, error) {
	conn := dbConn.GetSqlConn()
	var user mysql_model.User
	if conn == nil {
		log.Println("[Service Register] Get Sql connection error")
		return user, errors.New("something error")
	}
	//生成16位的盐值
	salt := utils.RandomStringRune(16)
	hash := sha1.New()
	_, err := hash.Write([]byte(pwd + salt))
	bs := hex.EncodeToString(hash.Sum(nil))
	if err != nil {
		return user, err
	}
	//加密后的密码=salt:hash(pwd+salt)
	password := salt + ":" + bs
	//创建token
	token := utils.GenToken(telephone)
	if err := conn.UserSignUp(telephone, password, token); err != nil {
		return user, err
	}

	user, err = conn.GetUserInfo(telephone)
	if err != nil {
		return user, err
	}

	return user, nil
}
