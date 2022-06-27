package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName   string
	Email      string
	Phone      string
	ClientCode string
	RoleId     uint
	AdminUser  bool
	Password   string
	UserStatus string
	Client     Client
	Role       Role
}
