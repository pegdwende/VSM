package models

import (
	AuthModels "github.com/pegdwende/VSM.git/auth/models"
	"github.com/pegdwende/VSM.git/database"
	"gorm.io/gorm"
)

type OrderItems struct {
	gorm.Model
	ProductID uint
	Product   Product
	OrderID   uint
	Order     Order
	Quantity  uint
	ClientID  uint
	Client    AuthModels.Client
}

func (OrderItems *OrderItems) Create() {
	database.GetConnection().Create(&OrderItems)
}
