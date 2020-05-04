//+build wireinject

package main

import (
	"github.com/knry0329/go-di/controller"
	"github.com/knry0329/go-di/repository"
	"github.com/knry0329/go-di/service"

	"github.com/google/wire"
)

func initializeUser() controller.UserController {
	wire.Build(
		repository.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return nil
}

func initializeFile() controller.FileController {
	wire.Build(
		repository.NewFileRepository,
		service.NewFileService,
		controller.NewFileController,
	)
	return nil
}
