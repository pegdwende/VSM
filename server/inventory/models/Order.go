package models

import (
	AuthModels "github.com/pegdwende/VSM.git/auth/models"
	"github.com/pegdwende/VSM.git/database"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID   uint
	User     AuthModels.User
	status   string
	ClientId uint
	Client   AuthModels.Client
}

func (Order *Order) create() {
	database.GetConnection().Create(&Order)
}
