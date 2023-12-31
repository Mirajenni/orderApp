package db

import (
	"fmt"
	"orderapp/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeConn() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Brazil/DeNoronha",
		config.Configuration.Database.DBHost,
		config.Configuration.Database.DBUser, config.Configuration.Database.DBPassword,
		config.Configuration.Database.DBName, config.Configuration.Database.DBPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&ProductDTO{})
}
