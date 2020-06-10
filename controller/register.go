package controller

import (
	"fmt"
	_ "im-server-go/docs"
	"im-server-go/domain"
	"im-server-go/entity/mysql_model"
	"im-server-go/service"
	"log"
	"regexp"
)

// @Summary  注册
// @Description 用户注册
// @Tags 测试
// @Accept mpfd
// @Produce json
// @Param telephone formData string true "手机号"
// @Param password formData string true "密码"
// @Success 200 {string} string "{"msg": "注册成功"}"
// @Failure 400 {string} string "{"msg": "注册失败"}"
// @Router /register [post]
func Register(c *domain.Context) {
	type reqBody struct {
		Telephone string `form:"telephone"`
		Password  string `form:"password"`
	}
	var reqbody reqBody

	err := c.ShouldBind(&reqbody)
	if err != nil {
		_ = c.AbortWithError(500, err)
		return
	}
	if len(reqbody.Telephone) != 11 || len(reqbody.Password) < 6 {
		log.Println("[Register]Telephone or password error")
		c.Fail("Telephone or password error")
		return
	}

	//校验手机号
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	compile := regexp.MustCompile(regular)
	if result := compile.MatchString(reqbody.Telephone); result == false {
		fmt.Println("[Register]Telephone number format is error")
		c.Fail("Telephone number format is error")
		return
	}
	var user mysql_model.User
	user, err = service.Register(reqbody.Telephone, reqbody.Password)
	if err != nil {
		c.Fail("This telephone has been registered")
		return
	}
	//设置AccessToken

	c.Header("AccessToken", user.AccessToken)
	c.Success(user)
}
