package models

import "gorm.io/gorm"

const (
	CREATE_ROLE = "create_role"
	CREATE_USER = "create_user"
)

type Permission struct {
	gorm.Model
	PermissionKey         string
	PermissionDescription string
	ClientID              uint
	RoleID                uint
	Role                  Role
}
