package controller

import (
	"im-server-go/domain"
	"im-server-go/service"
	"log"
	"regexp"
	"strconv"
)

// @Summary 登录
// @Description 用户登录
// @Tags 测试
// @Accept mpfd
// @Produce json
// @Param telephone formData string true "手机号"
// @Param password formData string true "密码"
// @Success 200 {string} string "{"msg": "登录成功"}"
// @Failure 400 {string} string "{"msg": "登录失败"}"
// @Router /login [post]
func Login(c *domain.Context) {
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
		log.Println("[Register]Telephone number format is error")
		c.Fail("Telephone number format is error")
		return
	}
	user, err := service.Login(reqbody.Telephone, reqbody.Password)
	if err != nil {
		log.Println("[Register]Login error")
		c.Fail(err.Error())
		return
	} else {
		c.Header("LoginUserId", strconv.Itoa(user.UserId))
		c.Success(user)
	}
}
