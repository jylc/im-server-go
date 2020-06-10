package controller

import (
	"im-server-go/domain"
	"im-server-go/service/auth"
	"im-server-go/service/telephone"
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
		Telephone string `form:"telephone" json:"telephone"`
		Password  string `form:"password" json:"password"`
	}

	var reqbody reqBody
	err := c.ShouldBind(&reqbody)
	if err != nil {
		c.Fail(err.Error())
		return
	}

	if valid := telephone.IsValid(reqbody.Telephone); !valid {
		c.Fail("请输入正确的手机号")
		return
	}

	user, err := auth.Login(reqbody.Telephone, reqbody.Password)
	if err != nil {
		c.Fail(err.Error())
		return
	}

	c.Success(user)
}
