package configuration

import (
	"fmt"
	"go-blog/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

// CreateDatabase method creates database connection and returns database object
func CreateDatabase() {
	// getting env variables
	DbHost := os.Getenv("DB_HOST")
	DbUsername := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASS")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	// creating error variable
	var err error

	// creating dsn string
	dsn := "host=" + DbHost + " user=" + DbUsername + " password=" + DbPassword + " dbname=" + DbName + " port=" + DbPort + " sslmode=disable TimeZone=CET"

	// opening connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error opening database!")
		return
	}
}

// CreateMigrations creating migrations for all models
func CreateMigrations() {
	// migration for User model
	userErr := DB.AutoMigrate(&model.User{})
	if userErr != nil {
		fmt.Println("Error migrating User table!")
		return
	}

	// migration for Task model
	taskErr := DB.AutoMigrate(&model.Task{})
	if taskErr != nil {
		fmt.Println("Error migrating Task table!")
	}
}
