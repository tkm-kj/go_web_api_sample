package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/tkm-kj/go_web_api_sample/internal/handler"
)

func Register(e *echo.Echo) {
	e.GET("/hello", handler.GetHello)
}
