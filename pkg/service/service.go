package service

import (
	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/repository"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) (bool, error)
	SetRole(user *models.User, role string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type Service struct {
	UserService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(repo.UserRepository),
	}
}
