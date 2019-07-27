package v1

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go_webapp/models"
	"go_webapp/pkg/code"
	"go_webapp/pkg/logging"
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

	valid := validation.Validation{}

	valid.Required(name, "name").Message("名称不为空")
	valid.Required(name, "createdBy").Message("创建人不为空")
	valid.MaxSize(name, 100, "createdBy").Message("创建人最大长度为100")
	valid.MaxSize(name, 100, "name").Message("名称最大长度为100")

	result_code := code.ERROR
	msg := ""
	result_info := false

	if !valid.HasErrors() {

		result_code = code.SUCCESS
		msg = "SUCCESS"
		result_info = models.AddTag(name, state, createdBy)

	} else {

		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
			msg = err.Message
		}

	}

	context.JSON(http.StatusOK, gin.H{
		"code": result_code,
		"msg":  msg,
		"data": result_info,
	})

}

func EditTag(context *gin.Context) {

	c := code.INVALID_PARAMS
	msg := "INVALID PARAMS"
	result := false

	id, _ := com.StrTo(context.Param("id")).Int()
	state, _ := com.StrTo(context.Query("state")).Int()

	name := context.Query("name")
	createdBy := context.Query("createdBy")

	if models.CheckExistsById(id) == true {
		c = code.SUCCESS
		msg = "SUCCESS"
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
		msg = "SUCCESS"
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
	c := code.SUCCESS
	msg := "SUCCESS"

	var data interface{}

	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("ID 必须大于0")

	if !valid.HasErrors() {
		data = models.GetTagById(id)
	} else {
		c = code.INVALID_PARAMS
		msg = "INVALID_PARAMS"
		data = false
	}

	context.JSON(http.StatusOK, gin.H{
		"code": c,
		"msg":  msg,
		"data": data,
	})

}
