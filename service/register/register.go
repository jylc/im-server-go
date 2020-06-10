package register

import (
	"errors"
	"im-server-go/connections"
	m "im-server-go/entity/mysql_model"
	"im-server-go/service/auth"
	"im-server-go/utils"
	"math/rand"
	"time"
)

func Register(telephone, pwd string) (*m.User, error) {
	db := connections.GetInstance()

	count := 0
	err := db.Where("telephone = ?", telephone).Count(&count).Error
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("此手机号已被注册")
	}

	salt := utils.GenSalt(16)
	user := &m.User{
		Telephone:     telephone,
		Password:      auth.GenPassword(pwd, salt),
		Nickname:      GenDefaultNick(),
		Sex:           0,
		Avatar:        "",
		Status:        1,
		AccessToken:   utils.GenToken(telephone),
		Ctime:         time.Now().Unix(),
		LastLoginTime: time.Now().Unix(),
	}
	err = db.Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GenDefaultNick() string {
	alpha := []rune("abcdefghijklmnopqrstuvwxyz")
	defaultNick := make([]rune, 10)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range defaultNick {
		defaultNick[i] = alpha[r.Intn(len(alpha))]
	}
	return string(defaultNick)
}
