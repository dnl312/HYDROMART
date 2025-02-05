package middleware

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	NewValidator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.NewValidator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	return nil
}
