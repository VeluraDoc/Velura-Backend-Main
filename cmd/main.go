package main

import (
	"log"

	"github.com/VeluraDoc/Velura-Backend-Main/internal/config"
	auth_handler "github.com/VeluraDoc/Velura-Backend-Main/internal/module/auth/handler"
	auth_middleware "github.com/VeluraDoc/Velura-Backend-Main/internal/module/auth/middleware"
	user_handler "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/handler"
	user_model "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/model"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.ConnectToDB()

	if err := db.AutoMigrate(&user_model.User{}); err != nil {
		log.Fatal("Error during migration: ", err)
	}

	port := "8080"
	if envPort := config.GetEnv("PORT"); envPort != "" {
		port = envPort
	}

	r := gin.Default()

	r.POST("/auth/sign-up", auth_handler.SignUp)
	r.POST("/auth/login", auth_handler.Login)

	protectedUserRoutes := r.Group("/user").Use(auth_middleware.AuthMiddleware())
	protectedUserRoutes.GET("/me", user_handler.GetMe)

	r.Run(":" + port)
}
