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

	name := context.Query("name")
	state, _ := com.StrTo(context.Query("state")).Int()
	createdBy := context.Query("createdBy")

	context.JSON(http.StatusOK, gin.H{
		"code": code.SUCCESS,
		"msg":  "SUCCESS",
		"date": models.AddTag(name, state, createdBy),
	})

}

func EditTag(context *gin.Context) {

	c := code.INVALID_PARAMS
	msg := "INVALID PARAMS"
	result := false

	id, _ := com.StrTo(context.Query("id")).Int()
	state, _ := com.StrTo(context.Query("state")).Int()

	name := context.Query("name")
	createdBy := context.Query("createdBy")

	if models.CheckExistsById(id) == true {
		c = code.SUCCESS
		result = models.EditTag(id, name, state, createdBy)
	}

	context.JSON(http.StatusOK, gin.H{
		"code": c,
		"msg":  msg,
		"data": result,
	})

}

func DeleteTag(context *gin.Context) {

	id, _ := com.StrTo(context.Param("id")).Int()

	c := code.INVALID_PARAMS
	msg := "INVALID PARAMS"
	result := false

	if models.CheckExistsById(id) == true {

		c = code.SUCCESS
		result = models.DeleteTagById(id)
	}

	context.JSON(http.StatusOK, gin.H{
		"code": c,
		"msg":  msg,
		"data": result,
	})

}

func GetTagById(context *gin.Context) {

	id, _ := com.StrTo(context.Param("id")).Int()

	context.JSON(http.StatusOK, gin.H{
		"code": code.SUCCESS,
		"msg":  "SUCCESS",
		"data": models.GetTagById(id),
	})

}
