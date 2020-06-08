package connections

import (
	"fmt"
	_ "im-server-go/config"

	"testing"
)

func TestMysqlConn(t *testing.T) {
	conn := GetSqlConn()
	if success := conn.UserSignIn("13141514481", "123212"); success == true {
		user, err := conn.GetUserInfo("13141514481")
		if err != nil {
			fmt.Println("Get user information failed!!!")
		} else {
			fmt.Println("success!")
			fmt.Println("user phone number :", user.Telephone,
				" password :", user.Password, " nickname :", user.Nickname)
		}
	} else {
		fmt.Println("fail!")
	}
}

func TestSignUp(t *testing.T) {
	conn := GetSqlConn()
	//手机号假设正确
	if succ := conn.UserSignUp("13141507497", "miao~"); succ == true {
		fmt.Println("register success!!!")
	} else {
		fmt.Println("register failed!!!")
	}
}
