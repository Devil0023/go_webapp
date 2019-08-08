package app

import (
	"github.com/gin-gonic/gin"
	"go_webapp/pkg/code"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(ErrorCode int, data interface{}) {

	g.C.JSON(http.StatusOK, gin.H{
		"code": ErrorCode,
		"msg":  code.GetMsg(ErrorCode),
		"data": data,
	})

	return
}
