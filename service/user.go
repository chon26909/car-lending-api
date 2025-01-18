package service

import "gorm.io/gorm"

type userService struct {
	db *gorm.DB
}

type IUserService interface {
	Login(username, password string) error
	Register(username, password string) error
}

func NewUserService(db *gorm.DB) IUserService {
	return &userService{db: db}
}

func (s *userService) Login(username, password string) error {
	return nil
}

func (s *userService) Register(username, password string) error {
	// Implement registration logic here
	return nil
}
