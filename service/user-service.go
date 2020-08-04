package service

import (
	"github.com/mayaramachado/invoice-api/entity"
	"github.com/mayaramachado/invoice-api/repository"
)

type UserService interface {
	Save(user entity.User) (entity.User, error)
	Login(username string, password string) bool
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}

func (service *userService) Login(email string, password string) bool {
	return true
}


func (service *userService) Save(user entity.User) (entity.User, error) {
	// criptografar a senha
	return service.userRepository.Save(user)
}
