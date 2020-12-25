package models

import (
	"gorm.io/gorm"
	"master/utils"
)

func init() {
	if utils.MysqlDBErr == nil && !utils.MysqlDB.Migrator().HasTable(&User{}) {
		_ = utils.MysqlDB.Migrator().CreateTable(&User{})
	}
}

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Nodes    []Node `gorm:"foreignKey:NodeID" json:"nodes"`
}

func AddUser(user User) {
	utils.MysqlDB.Create(&user)
}

func CheckUser(user User) bool {
	var count int64
	utils.MysqlDB.Model(&User{}).Where("Username = ?", user.Username).Count(&count)
	if count > 0 {
		return true
	}
	return false
}
