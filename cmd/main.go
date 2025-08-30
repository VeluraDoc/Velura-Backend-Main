package main

import (
	"log"

	"github.com/VeluraDoc/Velura-Backend-Main/config"
	"github.com/VeluraDoc/Velura-Backend-Main/internal/handler"
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

	r.POST("/auth/sign-up", handler.SignUpHandler)
	r.POST("/auth/login", handler.LoginHandler)

	r.Run(":" + port)
}
