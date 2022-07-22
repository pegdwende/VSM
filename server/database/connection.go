package database

import (
	"github.com/pegdwende/VSM.git/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var databaseConnectionEstablished = false

func Connect() {

	dns := Dns{
		username: env.GetRequiredEnvVariable("MYSQL_USER"),
		password: env.GetRequiredEnvVariable("MYSQL_PASSWORD"),
		host:     env.GetRequiredEnvVariable("MYSQL_DB_HOST"),
		port:     env.GetRequiredEnvVariable("MYSQL_DB_PORT"),
		protocol: env.GetRequiredEnvVariable("MYSQL_DB_PROTOCOL"),
		database: env.GetRequiredEnvVariable("MYSQL_DATABASE"),
	}

	connection, err := gorm.Open(mysql.Open(dns.getDnsString()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

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
