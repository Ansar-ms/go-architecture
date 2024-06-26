package repository

import (
	"errors"
	"go-architecture/pkg/model"
	"sync"
)

// UserRepository handles storage and retrieval of users
type UserRepository struct {
	mu    sync.Mutex
	users []model.User
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []model.User{},
	}
}

// GetAllUsers returns all users
func (r *UserRepository) GetAllUsers() []model.User {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.users
}

// AddUser adds a new user
func (r *UserRepository) AddUser(user model.User) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user.ID = len(r.users) + 1 // Assign ID (simple increment)
	r.users = append(r.users, user)
}

// UpdateUser updates an existing user by ID
func (r *UserRepository) UpdateUser(updatedUser model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for idx, user := range r.users {
		if user.ID == updatedUser.ID {
			r.users[idx] = updatedUser
			return nil
		}
	}
	return errors.New("user not found")
}

// DeleteUser deletes a user by ID
func (r *UserRepository) DeleteUser(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for idx, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:idx], r.users[idx+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *UserRepository) GetUserByName(name string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Name == name {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
