package services

import (
	"errors"
	"go-architecture/pkg/model"
	"go-architecture/pkg/repository"
)

// UserService provides user-related operations
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService creates a new UserService instance
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// GetAllUsers returns all users
func (s *UserService) GetAllUsers() []model.User {
	return s.repo.GetAllUsers()
}

// AddUser adds a new user
func (s *UserService) AddUser(user model.User) {
	s.repo.AddUser(user)
}

// UpdateUser updates an existing user by ID
func (s *UserService) UpdateUser(updatedUser model.User) error {
	return s.repo.UpdateUser(updatedUser)
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) AuthenticateUser(name, password string) (*model.User, error) {
	user, err := s.repo.GetUserByName(name)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
