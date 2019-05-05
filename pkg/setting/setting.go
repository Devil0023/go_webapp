package setting

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File
)

func init() {

	//Cfg, err = ini.load("conf/app." + env +".ini")
	fmt.Println(env)
}
