package models

import (
	"gorm.io/gorm"
)

const USER_ACTIVE_STATUS = "A"
const USER_TERMINATED_SATUS = "T"

type User struct {
	gorm.Model
	UserName   string `json:"username"`
	Email      string `json:"email" gorm:"unique"`
	Phone      string `json:"phone_number"`
	ClientCode string `json:"client_code"`
	ClientID   uint
	RoleId     uint   `json:"-"`
	AdminUser  bool   `json:"-"`
	Password   string `json:"-"`
	UserStatus string `json:"-"`
	Client     Client
	Role       Role `json:"-"`
	// CreatedAt  time.Time `gorm:"type:datetime" json:"created_at,string,omitempty"`
	// UpdatedAt  time.Time `gorm:"type:datetime" json:"updated_at,string,omitempty"`
	// DeletedAt  time.Time `gorm:"type:datetime" json:"deleted_at,string,omitempty"`
}
