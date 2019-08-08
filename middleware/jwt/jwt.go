package jwt

import (
	"github.com/gin-gonic/gin"
	"go_webapp/pkg/app"
	"go_webapp/pkg/code"
	"go_webapp/pkg/util"
	"time"
)

func JWT() gin.HandlerFunc {

	return func(context *gin.Context) {

		appG := app.Gin{context}
		data := make(map[string]interface{})

		token := context.Query("token")

		if token == "" {

			appG.Response(code.ERROR_AUTH, data)

			return

		} else {

			claims, err := util.ParseToken(token)

			if err != nil {
				appG.Response(code.ERROR_AUTH_CHECK_TOKEN_FAIL, data)
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				appG.Response(code.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, data)
				return
			}

		}

		context.Next()

	}
}
