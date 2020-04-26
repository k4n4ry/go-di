package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// func GetUser() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		id := c.Param("id")
// 		return c.String(http.StatusOK)
// 	}
// }

func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
