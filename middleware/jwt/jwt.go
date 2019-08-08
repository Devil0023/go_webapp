package jwt

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go_webapp/pkg/app"
	"go_webapp/pkg/code"
	"go_webapp/pkg/util"
	"time"
)

type JWTAuth struct {
	Token string `valid:"Required"`
}

func JWT() gin.HandlerFunc {

	return func(context *gin.Context) {

		appG := app.Gin{context}
		data := make(map[string]interface{})
		valid := validation.Validation{}

		auth := JWTAuth{Token: context.GetHeader("M-Token")}

		ok, _ := valid.Valid(&auth)

		if !ok {

			appG.Response(code.ERROR_AUTH, data)
			context.Abort()
			return

		} else {

			claims, err := util.ParseToken(auth.Token)

			if err != nil {
				appG.Response(code.ERROR_AUTH_CHECK_TOKEN_FAIL, data)
				context.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				appG.Response(code.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, data)
				context.Abort()
				return
			}

		}

		context.Next()

	}
}
