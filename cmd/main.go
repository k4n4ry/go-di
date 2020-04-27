package main

import (
	"net/http"

	"github.com/knry0329/go-di/db"

	"github.com/knry0329/go-di/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
書きをもとに、DIライクな実装をしてみる!
https://recruit-tech.co.jp/blog/2018/12/02/go_after_one_year_from_the_start/
・service, repositoryを作る
・DBをdockerで建てる
・mainからインジェクションするか、コンテナインスタンスを作ってみる(ライブラリがあれば、使ってみる)
*/
func main() {
	// db初期化
	if err := db.GormConnect("mysql", "knr:knrpw@tcp(localhost:3308)/knrdb?charset=utf8&parseTime=True&loc=Local"); err != nil {
		panic(err)
	}
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/users/:id", controller.GetUser)

	// Start server
	e.Logger.Fatal(e.Start(":1333"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
