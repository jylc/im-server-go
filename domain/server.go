package domain

import (
	"github.com/gin-gonic/gin"
	"im-server-go/utils"
	"log"
	"net/http"
	"time"
)

type MustParams struct {
	AccessToken string
	UserID      int
	IsLogin     bool
}

type Context struct {
	*gin.Context
	Param *MustParams
}

type Controller func(c *Context)

func NewCtrl(controller Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("AccessToken")
		userId := utils.AsInt(c.GetHeader("LoginUserId"))
		isLogin := false
		if userId != 0 {
			isLogin = true
		}
		params := &MustParams{
			AccessToken: accessToken,
			UserID:      userId,
			IsLogin:     isLogin,
		}
		ctx := &Context{c, params}
		controller(ctx)
	}
}

func (c *Context) Success(data ...interface{}) {
	c.JSON(http.StatusOK, success(data))
}

func (c *Context) Fail(message string, obj ...interface{}) {
	log.Println("req 500", "500:"+message)
	c.JSON(http.StatusOK, fail(message, obj))
}

func success(obj ...interface{}) interface{} {
	resp := gin.H{
		"code":      200,
		"data":      obj,
		"message":   "",
		"timestamp": time.Now().Unix(),
	}
	return resp
}

func fail(msg string, obj ...interface{}) interface{} {
	return gin.H{
		"code":      500,
		"data":      obj,
		"message":   msg,
		"timestamp": time.Now().Unix(),
	}
}
