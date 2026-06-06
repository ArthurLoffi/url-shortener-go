package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"url-shortener-go/internal/core/domain"
)

var Database *gorm.DB

func Connect() {
	// Connect with gorm
	err := godotenv.Load()
	if err != nil {
		log.Println(".ENV not found, using environment variables.")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get DB object: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	fmt.Println("Successfully connect to Neon Postgres Database!")

	err = db.AutoMigrate(&domain.User{}, &domain.Url{}, &domain.Click{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Variável global
	Database = db
	log.Println("Database migrated successfully!")
}