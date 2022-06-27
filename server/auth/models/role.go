package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleDescription string
	ClientCode      string
	Permissions     []Permission
}
