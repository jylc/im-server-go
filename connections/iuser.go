package connections

import (
	"im-server-go/entity/mysql_model"
)

//创建通用用户表操作接口
type IUser interface {
	//用户注册
	UserSignUp(telephone, pwd, token string) error
	//用户登录
	UserSignIn(telephone, pwd, token string) bool
	//获取用户信息
	GetUserInfo(telephone string) (mysql_model.User, error)
	//更新用户Token
	UpdateToken(user *mysql_model.User, token string) bool
}
