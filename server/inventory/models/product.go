package models

import (
	AuthModels "github.com/pegdwende/VSM.git/auth/models"
	"github.com/pegdwende/VSM.git/database"
	"gorm.io/gorm"
)

const (
	PENDING_ORDER   int32 = 1
	COMPLETED_ORDER int32 = 2
	CANCELLED_ORDER int32 = 3
)

type Product struct {
	gorm.Model
	ProductName        string
	ProductDescription string
	ProductImage       string
	ProductQuantity    uint
	ClientID           uint
	Client             AuthModels.Client
	ProductPrice       float64
}

func (Product *Product) Create() {
	database.GetConnection().Create(&Product)
}

func (Product *Product) Update() {
	database.GetConnection().Save(&Product)
}

func (Product *Product) Delete() {
	database.GetConnection().Delete(&Product)
}
