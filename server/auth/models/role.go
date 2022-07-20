package models

import (
	"github.com/pegdwende/VSM.git/database"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleDescription string
	ClientID        uint
	Client          Client
	Permissions     []Permission
}

func (role *Role) create() {
	database.GetConnection().Create(&role)
}
