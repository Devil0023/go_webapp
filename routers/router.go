package routers

import (
	"github.com/gin-gonic/gin"
	"go_webapp/controller/api"
	"go_webapp/controller/api/v1"
	"go_webapp/middleware/jwt"
	"go_webapp/pkg/setting"
)

//InitRouter
func InitRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	auth := router.Group("/auth")
	{
		auth.POST("", api.GetAuth)
	}

	apiv1 := router.Group("/api/v1")
	apiv1.Use(jwt.JWT())

	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.GET("/tags/:id", v1.GetTagById)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return router
}
