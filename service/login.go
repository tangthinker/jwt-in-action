package service

import (
	"errors"
	"jwt-in-action/auth"
)

type LoginService interface {
	Login(username, password string) (string, error)
	Register(username, password string) (string, error)
}

type loginService struct {
	userService UserService
}

func NewLoginService() LoginService {
	return &loginService{
		userService: NewUserService(),
	}
}

func (l *loginService) Login(username, password string) (string, error) {
	user, err := l.userService.GetByUsername(username)

	if err != nil {
		return "", err
	}

	if user.Password == password {
		token, err := auth.GenerateJWT(user.Name)
		if err != nil {
			return "", err
		}
		return token, nil
	}

	return "", errors.New("invalid username or password")

}

func (l *loginService) Register(username, password string) (string, error) {
	return "", errors.New("not implemented")
}
