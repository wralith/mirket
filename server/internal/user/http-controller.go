package user

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// TODO: Auth
// TODO: Errors

type HTTPController struct {
	svc *Service
}

func NewHTTPController(svc *Service) *HTTPController {
	return &HTTPController{svc: svc}
}

type RegisterRequest struct {
	Username string `json:"username" validate:"min=4"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=6"`
}

func (cnt *HTTPController) Register(c echo.Context) error {
	var request RegisterRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unknown error")
	}
	newUserOpts := NewUserOpts{Username: request.Username, Email: request.Email, HashedPassword: hash}

	if err := cnt.svc.Create(context.Background(), newUserOpts); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()) // TODO: More error checks
	}

	return c.NoContent(http.StatusCreated)
}

type LoginRequest struct {
	Username string `json:"username" validate:"min=4"`
	Password string `json:"password" validate:"min=6"`
}

func (cnt *HTTPController) Login(c echo.Context) error {
	var request LoginRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := cnt.svc.CheckPassword(context.Background(), request.Username, request.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid credentials")
	}
	user, err := cnt.svc.GetByUsername(context.Background(), request.Username)
	return c.JSON(http.StatusOK, user) // Send token or something.

}
