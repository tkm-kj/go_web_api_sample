package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tkm-kj/go_web_api_sample/internal/middleware"
	"github.com/tkm-kj/go_web_api_sample/internal/repository"
	"github.com/tkm-kj/go_web_api_sample/internal/serializer"
	"github.com/tkm-kj/go_web_api_sample/internal/usecase"
)

func GetPost(c echo.Context) error {
	u := usecase.NewPostUsecase(repository.NewPostRepository(middleware.GetMySQLConnection()))
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handleError(err, c)
	}
	post, err := u.Get(id)
	if err != nil {
		return handleError(err, c)
	}

	return c.JSON(http.StatusOK, serializer.NewPostSerializer(post))
}
