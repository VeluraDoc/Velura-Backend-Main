package user_router

import (
	auth_middleware "github.com/VeluraDoc/Velura-Backend-Main/internal/module/auth/middleware"
	user_handler "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	protected := r.Group("/")
	protected.Use(auth_middleware.AuthMiddleware())

	protected.GET("/me", user_handler.GetMe)
}
