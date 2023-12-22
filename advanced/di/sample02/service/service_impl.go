package service

import (
	"github.com/luismayta/learn-go/advanced/di/sample02/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (s *UserServiceImpl) GetUserName(userID int) string {
	return s.UserRepository.GetUserByID(userID)
}
