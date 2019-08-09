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
	Ini string
	Env string
)

type App struct {
	JwtSecret string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Log struct {
	LogSavePath   string
	LogSaveName   string
	LogFileExt    string
	LogTimeFormat string
}

var LogSetting = &Log{}

type Redis struct {
	Host        string
	Password    string
	Database    int
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

//Setup 注册配置文件
func Setup() {

	LoadEnv() //加载配置文件

	LoadServer()       //加载服务配置
	LoadDatabase()     //加载数据库配置
	LoadRedisSetting() //加载Redis配置
	LoadApp()          //加载应用配置
	LoadLogSetting()   //加载日志配置

}

//LoadEnv 加载环境设置
func LoadEnv() {

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
		log.Fatal("Failed to parse ini: ", err)
	}
}

//LoadDatabase 加载数据库配置
func LoadDatabase() {

	err := Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatal("Failed to map DatabaseSetting")
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

//LoadServer 加载服务器配置
func LoadServer() {
	err := Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatal("Failed to map ServerSetting")
	}
}

//LoadApp 加载应用配置
func LoadApp() {
	err := Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatal("Failed to map AppSetting")
	}
}

//LoadLogSetting 加载日志配置
func LoadLogSetting() {

	err := Cfg.Section("log").MapTo(LogSetting)
	if err != nil {
		log.Fatal("Failed to map LogSetting")
	}
}

//LoadRedisSetting 加载Redis配置
func LoadRedisSetting() {
	err := Cfg.Section("redis").MapTo(RedisSetting)

	if err != nil {
		log.Fatal("Failed to map RedisSetting")
	}

	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}
