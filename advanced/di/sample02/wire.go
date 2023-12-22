package sample02

import (
	"github.com/google/wire"

	"github.com/luismayta/learn-go/advanced/di/sample02/repository"
	"github.com/luismayta/learn-go/advanced/di/sample02/service"
)

func InitializeUserService() service.UserService {
	wire.Build(repository.NewUserRepository,
		service.NewUserServiceImpl)
	return nil
}
