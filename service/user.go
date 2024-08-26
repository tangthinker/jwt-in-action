package service

import "jwt-in-action/model"

type UserService interface {
	GetUser(id uint) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	GetByUsername(username string) (*model.User, error)
}

type userService struct {
	userModel *model.UserModel
}

func NewUserService() UserService {
	return &userService{
		userModel: model.NewUserModel(),
	}
}

func (u userService) GetByUsername(username string) (*model.User, error) {
	return u.userModel.GetByName(username)
}

func (u userService) GetUser(id uint) (*model.User, error) {
	return u.userModel.Get(id)
}

func (u userService) CreateUser(user *model.User) error {
	return u.userModel.Create(user)
}

func (u userService) UpdateUser(user *model.User) error {
	return u.userModel.Update(user)
}

func (u userService) DeleteUser(id uint) error {
	return u.userModel.Delete(id)
}
