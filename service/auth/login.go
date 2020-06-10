package auth

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"im-server-go/connections"
	m "im-server-go/entity/mysql_model"
	"im-server-go/utils"
	"strings"
)

//登录
func Login(telephone, password string) (*m.User, error) {
	user, err := GetUserInfo(telephone)
	if err != nil {
		return nil, err
	}

	if right := CheckPassword(password); !right {
		return nil, errors.New("密码错误")
	}

	//用户登录时创建token
	token := utils.GenToken(telephone)
	UpdateToken(user.UserId, token)

	return user, nil
}

//获取通过手机号查询用户信息
func GetUserInfo(telephone string) (*m.User, error) {
	db := connections.GetInstance()

	user := &m.User{}
	err := db.Where("telephone = ?", telephone).Find(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

//用户登录
func UpdateToken(userId int, token string) {
	db := connections.GetInstance()

	user := &m.User{}
	db.Where("user_id = ?", userId).First(user)

	user.AccessToken = token
	db.Save(user)
}

func GenPassword(pwd, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(pwd + salt))

	return salt + ":" + hex.EncodeToString(hash.Sum(nil))
}

func CheckPassword(pwd string) bool {
	ss := strings.Split(pwd, ":")
	if len(ss) == 0 {
		return false
	}
	shouldBe := GenPassword(pwd, ss[0])
	if pwd != shouldBe {
		return false
	}
	return true
}
