package models

import (
	"../pkg/setting"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int               `gorm:"primary_key" json:"id"`
	Created_at orm.DateTimeField `json:"created_at"`
	Updated_at orm.DateTimeField `json:"updated_at"`
}

func init() {

	var (
		err                                                error
		dbType, dbName, user, password, host, table_prefix string
	)

	section, err := setting.Cfg.GetSection("database")

	if err != nil {
		log.Fatal(" Failed to get section database : ", err)
	}

	dbType = section.Key("TYPE").String()
	dbName = section.Key("NAME").String()
	user = section.Key("USER").String()
	password = section.Key("PASSWORD").String()
	host = section.Key("HOST").String()
	table_prefix = section.Key("TABLE_PREFIX").String()

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	fmt.Println(11111111)

	if db == nil {
		fmt.Println(22222222)
	} else {
		fmt.Println(33333333)
	}

	if err != nil {
		log.Fatal("Failed to connect database : ", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return table_prefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
