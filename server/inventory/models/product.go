package models

import (
	AuthModels "github.com/pegdwende/VSM.git/auth/models"
	"github.com/pegdwende/VSM.git/database"
	"gorm.io/gorm"
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
