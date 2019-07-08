package jwt

import (
	"github.com/gin-gonic/gin"
	"go_webapp/pkg/code"
	"go_webapp/pkg/util"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {

	return func(context *gin.Context) {
		var c int
		var data interface{}

		c = code.SUCCESS
		token := context.Query("token")

		if token == "" {
			c = code.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)

			if err != nil {
				c = code.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				c = code.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}

		}

		if c != code.SUCCESS {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": c,
				"msg":  "Auth Error",
				"data": data,
			})

			context.Abort()
			return
		}

		context.Next()

	}
}
