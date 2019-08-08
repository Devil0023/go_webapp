package v1

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go_webapp/models"
	"go_webapp/pkg/app"
	"go_webapp/pkg/code"
	"go_webapp/pkg/util"
)

func GetTags(context *gin.Context) {

	appG := app.Gin{context}
	data := make(map[string]interface{})

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

	data["total"] = models.GetTagTotal(maps)
	data["list"] = models.GetTags(page, pagesize, maps)

	appG.Response(code.SUCCESS, data)

}

func AddTag(context *gin.Context) {

	appG := app.Gin{context}
	data := make(map[string]interface{})
	valid := validation.Validation{}

	name := context.Query("name")
	state, _ := com.StrTo(context.Query("state")).Int()
	createdBy := context.Query("createdBy")

	valid.Required(name, "name").Message("名称不为空")
	valid.Required(name, "createdBy").Message("创建人不为空")
	valid.MaxSize(name, 100, "createdBy").Message("创建人最大长度为100")
	valid.MaxSize(name, 100, "name").Message("名称最大长度为100")

	if !valid.HasErrors() {

		data["result"] = models.AddTag(name, state, createdBy)
		appG.Response(code.SUCCESS, data)

	} else {

		util.LogErrors(valid.Errors)
		appG.Response(code.INVALID_PARAMS, data)
	}

	return
}

func EditTag(context *gin.Context) {

	appG := app.Gin{context}
	data := make(map[string]interface{})
	valid := validation.Validation{}

	id, _ := com.StrTo(context.Param("id")).Int()
	state, _ := com.StrTo(context.Query("state")).Int()
	name := context.Query("name")
	createdBy := context.Query("createdBy")

	valid.Required(name, "name").Message("名称不为空")
	valid.Required(name, "createdBy").Message("创建人不为空")
	valid.MaxSize(name, 100, "createdBy").Message("创建人最大长度为100")
	valid.MaxSize(name, 100, "name").Message("名称最大长度为100")

	if models.CheckExistsById(id) == true {
		data["result"] = models.EditTag(id, name, state, createdBy)
		appG.Response(code.SUCCESS, data)
	} else {
		appG.Response(code.INVALID_PARAMS, data)
	}

	return

}

func DeleteTag(context *gin.Context) {

	appG := app.Gin{context}
	data := make(map[string]interface{})

	id, _ := com.StrTo(context.Param("id")).Int()

	if models.CheckExistsById(id) == true {
		data["result"] = models.DeleteTagById(id)
		appG.Response(code.SUCCESS, data)
	} else {
		appG.Response(code.ERROR, data)
	}

	return

}

func GetTagById(context *gin.Context) {

	appG := app.Gin{context}
	data := make(map[string]interface{})
	valid := validation.Validation{}

	id, _ := com.StrTo(context.Param("id")).Int()

	valid.Min(id, 1, "id").Message("ID 必须大于0")

	if !valid.HasErrors() {
		data["result"] = models.GetTagById(id)
		appG.Response(code.SUCCESS, data)
	} else {
		util.LogErrors(valid.Errors)
		appG.Response(code.ERROR, data)
	}

	return

}
