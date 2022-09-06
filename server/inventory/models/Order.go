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
	Status   int32
	ClientID uint
	Client   AuthModels.Client
}

func (Order *Order) Create() {
	database.GetConnection().Create(&Order)
}
