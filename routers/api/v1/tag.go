package v1

import (
	"../../../models"
	"../../../pkg/code"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTags(context *gin.Context) {

	var state, page, pagesize int = -1, 1, 20

	maps := make(map[string]interface{})

	if context.Query("name") != "" {
		maps["name"] = context.Query("name")
	}

	if context.Query("state") != "" {
		state, _ = com.StrTo(context.Query("state")).Int()
		maps["state"] = state
	}

	if context.Query("page") != "" {
		page, _ = com.StrTo(context.Query("page")).Int()
	}

	if context.Query("pagesize") != "" {
		pagesize, _ = com.StrTo(context.Query("pagesize")).Int()
	}

	context.JSON(http.StatusOK, gin.H{
		"code":  code.SUCCESS,
		"msg":   "SUCCESS",
		"list":  models.GetTags(page, pagesize, maps),
		"total": models.GetTagTotal(maps),
	})

}

func AddTag(context *gin.Context) {

}

func EditTag(context *gin.Context) {

}

func DeleteTag(context *gin.Context) {

}
