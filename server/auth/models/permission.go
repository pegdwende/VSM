package models

import "gorm.io/gorm"

const (
	CREATE_USER = "create_user"
)

type Permission struct {
	gorm.Model
	PermissionKey         string
	PermissionDescription string
	ClientCode            string
	RoleId                uint
}