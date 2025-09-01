// Copyright 2025 Velura Team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"

	"github.com/VeluraDoc/Velura-Backend-Main/internal/config"
	auth_router "github.com/VeluraDoc/Velura-Backend-Main/internal/module/auth"
	user_router "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user"
	user_model "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/model"
	"github.com/gin-gonic/gin"

	_ "github.com/VeluraDoc/Velura-Backend-Main/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Velura Pdf-To-Word API Documentation
// @version         1.0
// @description     API documentation for Velura Pdf-To-Word.
// @BasePath        /

// @contact.name    Velura Team
// @contact.url     https://github.com/VeluraOpenSource

// @license.name    Apache 2.0
// @license.url     https://www.apache.org/licenses/LICENSE-2.0

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Bearer <token>
func main() {
	config.LoadEnv()
	gin.SetMode(config.GetEnv("GIN_MODE"))
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

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + port)
}
