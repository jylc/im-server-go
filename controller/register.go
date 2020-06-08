package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "im-server-go/docs"
	"im-server-go/domain"
	"im-server-go/service"
	"net/http"
	"regexp"
)

// @Summary 测试SayHello
// @Description 向你说Hello
// @Tags 测试
// @Accept mpfd
// @Produce json
// @Param who query string true "人名"
// @Success 200 {string} string "{"msg": "hello Razeen"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /hello [get]

func Register(c *gin.Context) {
	type reqBody struct {
		Telephone string `json:"telephone"`
		Pwd       string `json:"pwd"`
	}
	var reqbody reqBody
	/*
		绑定数据
	*/
	err := c.ShouldBind(&reqbody)
	if err != nil {
		_ = c.AbortWithError(500, err)
		return
	}
	if len(reqbody.Telephone) != 11 || len(reqbody.Pwd) < 6 {
		fmt.Println("[Register]Telephone number or password error")
		response := &domain.ApiResponse{}
		c.Data(http.StatusInternalServerError,
			"application/json",
			response.Error("Telephone number or password error", nil).JsonBytes())
		return
	}

	//校验手机号
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	compile := regexp.MustCompile(regular)
	if result := compile.MatchString(reqbody.Telephone); result == false {
		fmt.Println("[Register]Telephone number format is error")
		response := &domain.ApiResponse{}
		c.Data(http.StatusInternalServerError,
			"application/json",
			response.Error("Telephone number format is error", nil).JsonBytes())
		return
	}
	response := service.Register(reqbody.Telephone, reqbody.Pwd)
	c.Data(http.StatusOK, "application/json", response.JsonBytes())
}
