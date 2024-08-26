package model

import (
	"gorm.io/gorm"
	"jwt-in-action/db"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Age      int
	Gender   string
	Email    string
	Password string `gorm:"not null"`
	Phone    string
}

func (User) TableName() string {
	return "tb_users"
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel() *UserModel {
	return &UserModel{db: db.GlobalDB}
}

func (m *UserModel) Create(user *User) error {
	return m.db.Create(user).Error
}

func (m *UserModel) Update(user *User) error {
	return m.db.Save(user).Error
}

func (m *UserModel) Delete(id uint) error {
	return m.db.Delete(&User{}, id).Error
}

func (m *UserModel) Get(id uint) (*User, error) {
	var user User
	err := m.db.First(&user, id).Error
	return &user, err
}

func (m *UserModel) GetByName(name string) (*User, error) {
	var user User
	err := m.db.Where("name = ?", name).First(&user).Error
	return &user, err
}

func (m *UserModel) List() ([]User, error) {
	var users []User
	err := m.db.Find(&users).Error
	return users, err
}
