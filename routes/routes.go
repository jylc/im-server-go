package routes

import (
	_ "im-server-go/docs"
)

// @host 127.0.0.1:8080
// @BasePath /api/v1
/*func InitRoutes(r *gin.Engine) {
	// 使用gin插件支持跨域请求
	r.Use(interceptor.CORS())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", controller.Register)
	}
}*/
