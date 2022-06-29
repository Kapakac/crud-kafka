package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
)

var DB *gorm.DB

func init() {
	basePath, _ := os.Getwd()

	if err := godotenv.Load(filepath.Join(basePath, ".env")); err != nil {
		log.Print("No .env file found")
	}
}

func ConnectionDatabase() *gorm.DB {
	host := os.Getenv("APP_DB_HOST")
	user := os.Getenv("APP_DB_USERNAME")
	password := os.Getenv("APP_DB_PASSWORD")
	port := os.Getenv("APP_DB_PORT")
	dbname := os.Getenv("APP_DB_NAME")

	var err error
	dns :=
		fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbname)
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect database")
	}

	return DB
}
