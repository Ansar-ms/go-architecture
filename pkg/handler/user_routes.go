// handlers/user_routes.go
package handler

import (
	"go-architecture/pkg/middleware"
	"go-architecture/pkg/services"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registers user-related routes with the router
func RegisterUserRoutes(router *gin.Engine, userService *services.UserService) {
	userHandler := NewUserHandler(userService)
	authMiddleware := middleware.BasicAuthMiddleware(userService)

	router.GET("/users", userHandler.GetAllUsers)
	router.POST("/users", userHandler.AddUser)
	router.PUT("/users/:id", authMiddleware, userHandler.UpdateUser)
	router.DELETE("/users/:id", authMiddleware, userHandler.DeleteUser)
}
