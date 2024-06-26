// handlers/user_handler.go
package handler

import (
	"go-architecture/pkg/middleware"
	"go-architecture/pkg/model"
	"go-architecture/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users := h.userService.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *UserHandler) AddUser(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	h.userService.AddUser(newUser)
	middleware.LogMessage(("User added"))
	c.JSON(http.StatusOK, gin.H{
		"users": h.userService.GetAllUsers(),
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedUser model.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	updatedUser.ID = id

	if err := h.userService.UpdateUser(updatedUser); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": h.userService.GetAllUsers(),
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.userService.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": h.userService.GetAllUsers(),
	})
}
