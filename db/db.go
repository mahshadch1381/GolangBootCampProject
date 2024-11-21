package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=1234 dbname=project port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent),})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
  
	// Configure connection pool settings
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to configure database: %v", err)
	}

	sqlDB.SetMaxIdleConns(5)                  // Minimum idle connections
	sqlDB.SetMaxOpenConns(20)                 // Maximum open connections
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Connection lifetime
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // Idle connection timeout


}
func GetDB() *gorm.DB {
    return DB
}