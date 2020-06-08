package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "im-server-go/docs"
	"im-server-go/routes"
)

var logo = "(♥◠‿◠)ﾉﾞ  CIM version0.1 启动成功   ლ(´ڡ`ლ)ﾞ  \n" +
	"   ____ ___ __  __ \n" +
	"  / ___|_ _|  \\/  |\n" +
	" | |    | || |\\/| |\n" +
	" | |___ | || |  | |\n" +
	"  \\____|___|_|  |_|\n" +
	"                   "

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://razeen.me

// @contact.name Razeen
// @contact.url https://razeen.me
// @contact.email me@razeen.me

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()

	routes.InitRoutes(r)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("[CIM] start failed!!!")
		return
	}
	fmt.Println("[CIM] start successful!!!")
	fmt.Println(logo)
}
