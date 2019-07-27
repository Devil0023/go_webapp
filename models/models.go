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
	Deleted_at time.Time `json:"deleted_at"`
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
	db.Callback().Delete().Replace("gorm:delete", deletedAtCallback)

}

func createdAtCallback(scope *gorm.Scope) {

	if !scope.HasError() {

		nowTime := time.Now()

		if createTimeField, ok := scope.FieldByName("created_at"); ok {

			if createTimeField.IsBlank {
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

		if _, ok := scope.Get("gorm:update_column"); !ok {
			_ = scope.SetColumn("updated_at", time.Now())
		}

	}
}

func deletedAtCallback(scope *gorm.Scope) {

	if !scope.HasError() {

		var extraOption, sql string

		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedAtField, hasDeletedAtField := scope.FieldByName("deleted_at")

		if !scope.Search.Unscoped && hasDeletedAtField {
			sql = fmt.Sprintf(
				"UPDATE $v SET %v = %v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedAtField.DBName),
				scope.AddToVars(time.Now()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)
		} else {
			sql = fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)
		}

		fmt.Println(sql)

		scope.Raw(sql).Exec()

	}
}

func addExtraSpaceIfExist(str string) string {

	if str != "" {
		return " " + str
	} else {
		return ""
	}

}

func CloseDB() {
	defer db.Close()
}
