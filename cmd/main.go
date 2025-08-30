package main

import (
	"log"

	"github.com/VeluraDoc/Velura-Backend-Main/config"
	"github.com/VeluraDoc/Velura-Backend-Main/internal/handlers"
	"github.com/VeluraDoc/Velura-Backend-Main/internal/model"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.ConnectToDB()

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("Error during migration: ", err)
	}

	port := "8080"
	if envPort := config.GetEnv("PORT"); envPort != "" {
		port = envPort
	}

	r := gin.Default()

	r.POST("/auth/sign-up", handlers.SignUpHandler)
	r.POST("/auth/login", handlers.LoginHandler)

	r.Run(":" + port)
}
