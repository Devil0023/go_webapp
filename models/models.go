package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go_webapp/pkg/setting"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int       `gorm:"primary_key" json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	//Deleted_at orm.DateTimeField `json:"deleted_at"`
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

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Fatal("Failed to connect database : ", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return table_prefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.Callback().Create().Replace("gorm:update_time_stamp", createdAtCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updatedAtCallback)

}

func createdAtCallback(scope *gorm.Scope) {

	if !scope.HasError() {

		nowTime := time.Now().Unix()

		if createTimeField, ok := scope.FieldByName("created_at"); ok {

			if createTimeField.IsBlank {
				fmt.Println(nowTime)
				fmt.Println(createTimeField.Field)
				_ = createTimeField.Set(nowTime)
			}
		}

		if updateTimeField, ok := scope.FieldByName("updated_at"); ok {
			if updateTimeField.IsBlank {
				_ = updateTimeField.Set(nowTime)
			}
		}
	}
}

func updatedAtCallback(scope *gorm.Scope) {

	if !scope.HasError() {
		nowTime := time.Now().Unix()

		if updateTimeField, ok := scope.FieldByName("updated_at"); ok {
			updateTimeField.Set(nowTime)
		}
	}
}

func CloseDB() {
	defer db.Close()
}
