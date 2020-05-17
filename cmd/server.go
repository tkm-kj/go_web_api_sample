package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	app_middleware "github.com/tkm-kj/go_web_api_sample/internal/middleware"
	"github.com/tkm-kj/go_web_api_sample/internal/routes"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.Register(e)

	e.Logger.Fatal(e.Start(":1323"))

	defer app_middleware.CloseMySQLConnection()
}
