package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tkm-kj/go_web_api_sample/internal/middleware"
	"github.com/tkm-kj/go_web_api_sample/internal/repository"
	"github.com/tkm-kj/go_web_api_sample/internal/usecase"
)

func GetPost(c echo.Context) error {
	u := usecase.NewPostUsecase(repository.NewPostRepository(middleware.GetMySQLConnection()))
	// TODO: error handling
	id, _ := strconv.Atoi(c.Param("id"))
	post := u.Get(id)

	return c.String(http.StatusOK, fmt.Sprintf("post: %+v", post))
}
