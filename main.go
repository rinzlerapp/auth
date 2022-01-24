package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)


type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
  if err := cv.validator.Struct(i); err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return nil
}

func main() {
  e := echo.New()
  e.Validator = &CustomValidator{validator: validator.New()}

  e.GET("/health", func(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
  })

  e.Logger.Fatal(e.Start(":1234"))
}