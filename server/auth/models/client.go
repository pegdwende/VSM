package models

import (
	//"database/sql/driver"

	"gorm.io/gorm"
)

// type clientStatus string

// const (
// 	T clientStatus = "Terminated"
// 	A clientStatus = "Active"
// 	P clientStatus = "Pending"
// )

// func (cs *clientStatus) Scan(value interface{}) error {
// 	*cs = clientStatus(value.([]byte))
// 	return nil
// }

// func (cs clientStatus) Value() (driver.Value, error) {
// 	return string(cs), nil
// }

// type businessType string

// const (
// 	RETAIL businessType = "Retail"
// )

// func (bt *businessType) Scan(value interface{}) error {
// 	*bt = businessType(value.([]byte))
// 	return nil
// }

// func (bt businessType) Value() (driver.Value, error) {
// 	return string(bt), nil
// }

// type Client struct {
// 	gorm.Model
// 	ClientCode string `gorm:"primaryKey"`
// 	name       string
// 	address    string
// 	email      string `gorm:"unique"`
// 	ClientStatus clientStatus `sql:"type:client_status"`
// 	BusinessType businessType `sql:"type:client_status"`
// }
const (
	T string = "Terminated"
	A string = "Active"
	P string = "Pending"
)

const (
	RETAIL string = "Retail"
)

type Client struct {
	gorm.Model
	Code         string `gorm:"primaryKey"`
	Name         string
	Address      string
	Email        string `gorm:"unique"`
	ClientStatus string `json:"-"`
	BusinessType string
}
