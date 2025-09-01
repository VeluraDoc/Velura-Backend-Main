package main

import (
	"log"

	"github.com/VeluraDoc/Velura-Backend-Main/internal/config"
	auth_router "github.com/VeluraDoc/Velura-Backend-Main/internal/module/auth"
	user_router "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user"
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

	auth_router.RegisterRoutes(r.Group("/auth"))
	user_router.RegisterRoutes(r.Group("/user"))

	r.Run(":" + port)
}
