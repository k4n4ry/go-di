package controller

import (
	"net/http"
	"strconv"

	"github.com/knry0329/go-di/service"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	GetUser(c echo.Context) error
}

type userController struct {
	svc service.UserService
}

// var instance userController

func NewUserController(svc service.UserService) UserController {
	// instance = userController{
	// 	svc: svc,
	// }
	return &userController{
		svc: svc,
	}
}

func (ctl *userController) GetUser(c echo.Context) error {
	// func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	name, _ := ctl.svc.GetUserName(idInt)
	// 一旦こういうふうに実装しているけど、controllerがServiceを直接呼び出す

	return c.String(http.StatusOK, name)
}
