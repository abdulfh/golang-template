package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseOpen ..
func DatabaseOpen() (*gorm.DB, error) {
	// Local
	dbuser := "newuser"
	dbpass := "newuser"
	dbname := "golangbe"
	dbaddres := "localhost"

	dbport := "5432"
	sslmode := "disable"
	dbtimeout := "5"

	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%s", dbaddres, dbport, dbuser, dbpass, dbname, sslmode, dbtimeout)
	db, err := gorm.Open(postgres.Open(args), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	// db.AutoMigrate(&models.Book{})

	if err != nil {
		fmt.Println("Error Connecting Database : ", err)
		return db, err
	}

	sqldb, _ := db.DB()
	sqldb.SetConnMaxLifetime(10)
	sqldb.SetMaxIdleConns(25)
	sqldb.SetMaxOpenConns(10)

	return db, nil
}
