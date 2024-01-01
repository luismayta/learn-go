// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package sample02

import (
	"github.com/luismayta/learn-go/advanced/di/sample02/repository"
	"github.com/luismayta/learn-go/advanced/di/sample02/service"
)

// Injectors from wire.go:

func InitializeUserService() *service.UserServiceImpl {
	userRepository := repository.NewUserRepository()
	userServiceImpl := service.NewUserServiceImpl(userRepository)
	return userServiceImpl
}