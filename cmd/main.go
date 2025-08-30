package main

import (
	"github.com/VeluraDoc/Velura-Backend-Main/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	port := "8080"
	if envPort := config.GetEnv("PORT"); envPort != "" {
		port = envPort
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":" + port)

}
