package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Str2Md5(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
