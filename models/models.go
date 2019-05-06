package models

import (
	"../pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func init() {

	var (
		err                                        error
		dbName, user, password, host, table_prefix string
	)

	section, err := setting.Cfg.GetSection("database")

	if err != nil {
		log.Fatal(2, "Failed to get section database : %v", err)
	}

	dbName = section.Key("NAME").String()
	user = section.Key("USER").String()
	password = section.Key("PASSWORD").String()
	table_prefix = section.Key("TABLE_PREFIX").String()

	db, err := gorm.Open(dbName, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Fatal(2, "Failed to connect database : %v", err)
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
