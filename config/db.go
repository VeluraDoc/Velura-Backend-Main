package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {
	if DB != nil {
		return DB
	}

	dbUser := GetEnv("DB_USER")
	dbPassword := GetEnv("DB_PASSWORD")
	dbName := GetEnv("DB_NAME")
	dbHost := GetEnv("DB_HOST")
	dbPort := GetEnv("DB_PORT")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	DB = db

	fmt.Println("Successfully connected to the database!")

	return DB
}
