package internal

import (
	"github.com/Crunchies1/creatures_backend/models"
)

type UserService struct {
	db *models.Client
}

func NewUserService(client *models.Client) *UserService {
	return &UserService{
		db: client,
	}
}

// CreateUser handles user creation business logic
func (s *UserService) CreateUser(user *models.User) error {
	// Add any business logic here (validation, transformation, etc.)
	return s.db.CreateUser(user)
}

// GetUsers retrieves all users from the database
func (s *UserService) GetUsers() ([]models.User, error) {
	return s.db.GetUsers()
}

// GetUserByID retrieves a specific user by ID
func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.db.GetUserByID(id)
}
