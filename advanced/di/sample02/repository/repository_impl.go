package repository

import (
	"fmt"
)

type UserRepositoryImpl struct{}

func (r *UserRepositoryImpl) GetUserByID(userID int) string {
	return fmt.Sprintf("User %d", userID)
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
