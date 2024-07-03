package model

// User struct to hold user information
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

// CreateUserRequest struct to hold user creation information (without ID)
type CreateUserRequest struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}
