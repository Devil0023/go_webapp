package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/validation"
	"go_webapp/pkg/logging"
	"math/rand"
	"sort"
	"strings"
	"time"
)

//Str2Md5 String转MD5
func Str2Md5(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//LogErrors 记录错误日志
func LogErrors(errors []*validation.Error) {

	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}

type KsortResult map[string]string

//Ksort
func Ksort(params map[string]string) (keys []string, result map[string]string) {

	for key, _ := range params {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	result = make(map[string]string)

	for _, key := range keys {
		result[key] = params[key]
	}

	return keys, result
}

//Implode
func Implode(glue string, params []string) string {

	var str string

	for _, value := range params {
		str += value + glue
	}

	str = str[0 : len(str)-1]

	return str
}

// GenValidateCode 生成随机数
func GenValidateCode(width int) string {

	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder

	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}

	return sb.String()

}
