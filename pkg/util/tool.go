package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/validation"
	"go_webapp/pkg/logging"
)

func Str2Md5(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func LogErrors(errors []*validation.Error) {

	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
