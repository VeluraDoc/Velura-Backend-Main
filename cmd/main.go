package main

import (
	"log"

	"github.com/VeluraDoc/Velura-Backend-Main/internal/config"
	"github.com/VeluraDoc/Velura-Backend-Main/internal/handler"
	"github.com/VeluraDoc/Velura-Backend-Main/internal/model"
	"github.com/VeluraDoc/Velura-Backend-Main/internal/module/auth"
	auth_middleware "github.com/VeluraDoc/Velura-Backend-Main/internal/module/auth/middleware"
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

	r.POST("/auth/sign-up", auth.SignUpHandler)
	r.POST("/auth/login", auth.LoginHandler)

	protectedUserRoutes := r.Group("/user").Use(auth_middleware.AuthMiddleware())
	protectedUserRoutes.GET("/me", handler.GetMe)

	r.Run(":" + port)
}
