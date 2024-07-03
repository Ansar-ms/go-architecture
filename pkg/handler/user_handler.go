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

// UserResponse defines the structure of the response for user-related operations
type UserResponse struct {
	Users []model.User `json:"users"`
}

// ErrorResponse defines the structure of an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} UserResponse
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users := h.userService.GetAllUsers()
	c.JSON(http.StatusOK, UserResponse{Users: users})
}

// AddUser godoc
// @Summary Create a new user
// @Description Add a new user
// @Tags users
// @Accept json
// @Produce json
// @Param body body model.CreateUserRequest true "User data"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Router /users [post]
func (h *UserHandler) AddUser(c *gin.Context) {
	var newUser model.CreateUserRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}

	// Convert CreateUserRequest to User and set ID
	user := model.User{
		Name:     newUser.Name,
		Age:      newUser.Age,
		Password: newUser.Password,
	}

	h.userService.AddUser(user)
	middleware.LogMessage("User added")
	c.JSON(http.StatusOK, UserResponse{Users: h.userService.GetAllUsers()})
}

// UpdateUser godoc
// @Summary Update a user by ID
// @Description Update an existing user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param body body model.User true "Updated user data"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @securityDefinitions.apiKey Authorization
// @in header
// @name Authorization
// @Security JWT
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}

	var updatedUser model.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}
	updatedUser.ID = id

	if err := h.userService.UpdateUser(updatedUser); err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{Users: h.userService.GetAllUsers()})
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete an existing user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @securityDefinitions.apiKey Authorization
// @in header
// @name Authorization
// @Security JWT
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}

	if err := h.userService.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{Users: h.userService.GetAllUsers()})
}
