package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go_webapp/models"
	"go_webapp/pkg/code"
	"go_webapp/pkg/logging"
	"go_webapp/pkg/util"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(context *gin.Context) {

	username := context.Query("username")
	password := context.Query("password")

	valid := validation.Validation{}

	a := auth{Username: username, Password: password}

	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})

	c := code.INVALID_PARAMS

	if ok {

		isExist := models.CheckAuth(username, password)

		if isExist {

			token, err := util.GenerateToken(username, password)

			if err != nil {
				c = code.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				c = code.SUCCESS
			}

		} else {
			c = code.ERROR_AUTH
		}

	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code": c,
		"msg":  "MSG",
		"data": data,
	})

}
