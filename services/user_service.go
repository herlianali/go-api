package services

import (
	"go-api/models"
	"go-api/repositories"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.FindAll()
}

func (s *userService) GetUserByID(id uint) (models.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	return s.userRepository.Create(user)
}

func (s *userService) UpdateUser(user models.User) (models.User, error) {
	return s.userRepository.Update(user)
}

func (s *userService) DeleteUser(user models.User) error {
	return s.userRepository.Delete(user)
}
