package models

import (
	"gorm.io/gorm"
	"master/utils"
)

func init() {
	if utils.MysqlDBErr == nil && !utils.MysqlDB.Migrator().HasTable(&Node{}) {
		_ = utils.MysqlDB.Migrator().CreateTable(&Node{})
	}
}

type Node struct {
	gorm.Model
	IP          string `gorm:"not null;unique"`
	HostName    string `gorm:"not null;unique"`
	SSHUsername string `gorm:"not null"`
	SSHPassword string `gorm:"not null"`
}
