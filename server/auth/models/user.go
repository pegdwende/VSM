package models

import "gorm.io/gorm"

const USER_ACTIVE_STATUS = "A"
const USER_TERMINATED_SATUS = "T"

type User struct {
	gorm.Model
	UserName   string `json:"username"`
	Email      string `json:"email" gorm:"unique"`
	Phone      string `json:"phone_number"`
	ClientCode string `json:"client_code"`
	RoleId     uint   `json:"-"`
	AdminUser  bool   `json:"-"`
	Password   string `json:"-"`
	UserStatus string `json:"-"`
	Client     Client
	Role       Role `json:"-"`
}
