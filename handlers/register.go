package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func Register(c echo.Context) (err error) {
	u := new(RegisterRequest)

	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "OK")
}
