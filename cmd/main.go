package main

import (
	"net/http"
	"os"

	"github.com/knry0329/go-di/config"

	"github.com/knry0329/go-di/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
書きをもとに、DIライクな実装をしてみる!
https://recruit-tech.co.jp/blog/2018/12/02/go_after_one_year_from_the_start/
・service, repositoryを作る
・DBをdockerで建てる
・mainからインジェクションするか、コンテナインスタンスを作ってみる(ライブラリがあれば、使ってみる)
・digってやつが良さそう?
*/

/*
・テスト書いてみて、DIの良さを確認する。
・DIコンテナのライブラリ使ってみる。
*/
func main() {
	// db初期化
	if err := db.GormConnect(config.DbConfig{
		Dbms: "mysql",
		Dsn:  "knr:knrpw@tcp(localhost:3308)/knrdb?charset=utf8&parseTime=True&loc=Local",
	}); err != nil {
		panic(err)
	}

	// constructor injection
	// r := repository.NewUserRepository()
	// s := service.NewUserService(r)
	// c := controller.NewUserController(s)

	// c := initializeUser()

	// Echo instance
	e := echo.New()

	// lf (os.File型)はWriteメソッド持ってるので、io.Writerを実装している→LoggerConfigのOutputにわたすことができる。
	lf, _ := os.OpenFile("../test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer lf.Close()
	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: lf,
	}))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/users/:id", initializeUser().GetUser)
	e.GET("/file/:bucket", initializeFile().GetFile)

	// Start server
	e.Logger.Fatal(e.Start(":1333"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
