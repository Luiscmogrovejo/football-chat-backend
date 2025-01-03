package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yourusername/football-chat-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Config struct {
	DB     *gorm.DB
	Secret string
}

func InitConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found. Falling back to system environment variables.")
	} else {
		log.Println(".env file loaded successfully")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	jwtSecret := os.Getenv("JWT_SECRET")

	log.Printf("DB_HOST: %s, DB_PORT: %s, DB_USER: %s, DB_PASSWORD: %s, DB_NAME: %s, JWT_SECRET: %s",
		host, port, user, password, dbName, jwtSecret)

	requiredEnv := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "JWT_SECRET"}
	for _, key := range requiredEnv {
		if os.Getenv(key) == "" {
			log.Fatalf("Environment variable %s is required but not set", key)
		}
	}

	Config.Secret = jwtSecret

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	var err error
	Config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connection established successfully")

	Config.DB.AutoMigrate(
		&models.User{},
		&models.Match{},
		&models.Message{},
		&models.Prediction{},
		&models.Room{},
		&models.Stream{},
	)
}
