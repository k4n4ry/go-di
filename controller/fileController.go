package controller

import (
	"net/http"

	"github.com/knry0329/go-di/service"

	"github.com/labstack/echo/v4"
)

type FileController interface {
	GetFile(c echo.Context) error
}

type fileController struct {
	svc service.FileService
}

// var instance fileController

func NewFileController(svc service.FileService) FileController {
	// instance = fileController{
	// 	svc: svc,
	// }
	return &fileController{
		svc: svc,
	}
}

func (ctl *fileController) GetFile(c echo.Context) error {
	bucket := c.Param("bucket")

	str, _ := ctl.svc.GetFileName(bucket)

	return c.String(http.StatusOK, str)
}
