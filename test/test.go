package main

import "fmt"
import (
	"go_webapp/pkg/util"
)

func main() {
	string := "123456"
	fmt.Print(util.Str2Md5(string))
}
