package routes

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"im-server-go/controller"
	_ "im-server-go/docs"
	"im-server-go/interceptor"
)

// @host 127.0.0.1:8080
// @BasePath /api/v1
func InitRoutes(r *gin.Engine) {
	r.Use(interceptor.CORS())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", controller.Register)
	}
}
