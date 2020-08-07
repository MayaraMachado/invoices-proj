package service

import (
	base64 "encoding/base64"
	"errors"
	"regexp"

	"github.com/mayaramachado/invoice-api/entity"
	"github.com/mayaramachado/invoice-api/repository"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

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
	user, err := service.userRepository.GetByEmail(email)
	if err != nil {
		return false
	}

	// comparar senhas
	pwdEncrypt := base64.StdEncoding.EncodeToString([]byte(password))
	return user.Password == pwdEncrypt
}

func (service *userService) Save(user entity.User) (entity.User, error) {
	//validar email
	if (len(user.Email) < 5 && len(user.Email) > 254) || !emailRegex.MatchString(user.Email) {
		return user, errors.New("email inválido")
	}

	// criptografar a senha
	pwdEncrypt := base64.StdEncoding.EncodeToString([]byte(user.Password))
	user.Password = pwdEncrypt
	return service.userRepository.Save(user)
}
