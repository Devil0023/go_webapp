package routers

import "github.com/gin-gonic/gin"
import "../pkg/setting"

func InitRouter() *gin.Engine {
	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	route.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "test",
		})
	})

	return route
}
