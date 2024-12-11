package database

import (
	"Trainify/model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func Connect() {
	// Get DSN from environment variable
	dsn := os.Getenv("dsn")
	if dsn == "" {
		log.Fatal("Environment variable `dsn` is not set.")
	}
	log.Printf("Using DSN: %s", dsn)

	// Open connection to the MySQL database using Gorm
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection successful.")

	// Run migrations
	err = db.AutoMigrate(new(model.User))
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Assign database connection to global variable
	DBConn = db
}
