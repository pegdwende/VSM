package models

type RolePermission struct {
	Id            uint
	ClientCode    string
	PermissionKey string
	RoleId        uint
}
