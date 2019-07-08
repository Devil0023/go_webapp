package models

import (
	"go_webapp/pkg/util"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json: "id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth

	db.Select("id").Where(Auth{Username: username, Password: util.Str2Md5(password)}).First(&auth)

	if auth.ID > 0 {
		return true
	} else {
		return false
	}
}
