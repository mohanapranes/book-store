package config

import (
	"fmt"
	"log"
	"time"

	"github.com/mohanapranes/book-store/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect initializes the database connection with retries and environment variable configuration
func Connect() {
	log.Println("Establishing database connection...")

	// Use environment variables to make DSN configurable
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		utils.GetEnv("DB_HOST", "localhost"),
		utils.GetEnv("DB_USER", "postgres"),
		utils.GetEnv("DB_PASSWORD", "postgres"),
		utils.GetEnv("DB_NAME", "book-store"),
		utils.GetEnv("DB_PORT", "5432"),
		utils.GetEnv("DB_TIMEZONE", "Asia/Shanghai"),
	)

	// Attempt connection with retries
	var err error
	for i := 0; i < 3; i++ { // 3 retries
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			// Test the connection
			sqlDB, _ := db.DB()
			var pingErr error
			if pingErr = sqlDB.Ping(); pingErr == nil {
				log.Println("Database connected successfully")
				return
			}
			err = pingErr
		}

		log.Printf("Failed to connect to database (attempt %d): %v", i+1, err)
		time.Sleep(2 * time.Second) // Retry delay
	}

	// If all retries fail, log and exit
	log.Fatalf("Could not connect to the database: %v", err)
}

// GetDB returns a pointer to the database instance
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}
	return db
}
