package database

import (
	"fmt"

	"github.com/pegdwende/VSM.git/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var databaseConnectionEstablished = false

func Connect() {

	fmt.Printf("database %s", env.GetRequiredEnvVariable("VSM_DB_PASSWORD"))
	dns := Dns{
		username: env.GetRequiredEnvVariable("VSM_DB_USERNAME"),
		password: env.GetRequiredEnvVariable("VSM_DB_PASSWORD"),
		host:     env.GetRequiredEnvVariable("VSM_DB_HOST"),
		port:     env.GetRequiredEnvVariable("VSM_DB_PORT"),
		protocol: env.GetRequiredEnvVariable("VSM_DB_PROTOCOL"),
		database: env.GetRequiredEnvVariable("VSM_DB_DATABASE"),
	}

	connection, err := gorm.Open(mysql.Open(dns.getDnsString()), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}
	databaseConnectionEstablished = true

	db = connection
}

func GetConnection() *gorm.DB {

	if !IsDatabaseConnectionEstablished() {
		Connect()
	}

	return db

}

func Migrate(models ...interface{}) error {

	return db.AutoMigrate(models...)
}

func IsDatabaseConnectionEstablished() bool {

	return databaseConnectionEstablished
}
