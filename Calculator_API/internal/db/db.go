package db

import (
	"Calculator_API/internal/calculationService"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) { 
	var err error

	loadEnv()

	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "postgres")
	sslmode := getEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(&calculationService.Calculation{}); err != nil {
		log.Fatalf("Could not migrate %v", err)
	}

	return db, nil
}

func loadEnv() {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Join(filepath.Dir(filename), "..", "..")
	envPath := filepath.Join(dir, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Warning: .env file not found at %s: %v", envPath, err)
		
		alternativePaths := []string{
			".env",
			"../.env",
			"../../.env",
			"Calculator_API/.env",
		}
		
		for _, path := range alternativePaths {
			if err := godotenv.Load(path); err == nil {
				log.Printf("Loaded .env from: %s", path)
				return
			}
		}
		
		log.Println("No .env file found, using environment variables or defaults")
	} else {
		log.Printf("Loaded .env from: %s", envPath)
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}