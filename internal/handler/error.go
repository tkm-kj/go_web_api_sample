package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tkm-kj/go_web_api_sample/internal/repository"
	"github.com/tkm-kj/go_web_api_sample/internal/serializer"
)

type InternalServerError struct {
	err error
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("internal server error: %+v", e.err)
}

func handleError(err error, c echo.Context) error {
	switch err.(type) {
	case *repository.RecordNotFoundError:
		return c.JSON(http.StatusNotFound, serializer.NewErrorSerializer(err))
	default:
		e := &InternalServerError{err}
		return c.JSON(http.StatusInternalServerError, serializer.NewErrorSerializer(e))
	}
}
