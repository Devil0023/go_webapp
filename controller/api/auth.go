package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go_webapp/models"
	"go_webapp/pkg/app"
	"go_webapp/pkg/code"
	"go_webapp/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(context *gin.Context) {

	appG := app.Gin{context}

	username := context.Query("username")
	password := context.Query("password")

	valid := validation.Validation{}
	data := make(map[string]interface{})

	a := auth{Username: username, Password: password}

	ok, _ := valid.Valid(&a)

	if !ok {
		util.LogErrors(valid.Errors)
		appG.Response(code.INVALID_PARAMS, data)
		return
	}

	isExist := models.CheckAuth(username, password)

	if !isExist {
		appG.Response(code.ERROR_AUTH, data)
		return
	}

	token, err := util.GenerateToken(username, password)

	if err != nil {
		appG.Response(code.ERROR_AUTH_TOKEN, data)
		return
	}

	data["token"] = token

	appG.Response(code.SUCCESS, data)

	return
}
