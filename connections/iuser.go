package connections

import (
	"im-server-go/entity/mysql_model"
)

//创建通用用户表操作接口
type IUser interface {
	//用户注册
	UserSignUp(telephone, pwd string) bool
	//用户登录
	UserSignIn(telephone, pwd string) bool
	//获取用户信息
	GetUserInfo(telephone string) (mysql_model.User, error)
	//TODO 添加用户Token
}
