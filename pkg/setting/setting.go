package setting

import (
	"github.com/go-ini/ini"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	Cfg *ini.File

	Ini     string
	Env     string
	RunMode string

	HttpPort int

	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	JwtSecret string

	LogSavePath   string
	LogSaveName   string
	LogFileExt    string
	LogTimeFormat string
)

func init() {

	info, err := ioutil.ReadFile("./.env")

	if err != nil {
		log.Fatal("Failed to parse env: ", err)
	}

	Env = string(info[:])

	Ini = "conf/app." + Env + ".ini"

	_, err = os.Stat(Ini)

	if err != nil {
		Ini = "conf/app.ini"
	}

	Cfg, err = ini.Load(Ini)
	if err != nil {
		log.Fatal(" Failed to parse ini: ", err)
	}

	LoadRunMode()
	LoadHttpServer()
	LoadApp()
	loadLogSetting()
}

func LoadRunMode() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadHttpServer() {
	HttpPort = Cfg.Section("server").Key("HTTP_PORT").MustInt(80)
	ReadTimeout = time.Duration(Cfg.Section("server").Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(Cfg.Section("server").Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	JwtSecret = Cfg.Section("app").Key("JWT_SECRET").MustString("")
}

func loadLogSetting() {
	LogSavePath = Cfg.Section("log").Key("LOG_SAVE_PATH").MustString("")
	LogFileExt = Cfg.Section("log").Key("LOG_FILE_EXT").MustString("")
	LogSaveName = Cfg.Section("log").Key("LOG_SAVE_NAME").MustString("")
	LogTimeFormat = Cfg.Section("log").Key("TIME_FORMAT").MustString("")
}
