package routers

import "github.com/gin-gonic/gin"
import "../pkg/setting"
import "./api/v1"

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	apiv1 := router.Group("/api/v1")

	{
		apiv1.GET("/tags", func(context *gin.Context) {
			v1.GetTags(context)
		})
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return router
}
