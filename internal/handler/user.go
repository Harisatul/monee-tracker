package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"technopartner/internal/entity"
	"technopartner/internal/service/user"
)

type AuthHandler struct {
	authService user.AuthService
}

func New(authService user.AuthService) *AuthHandler {
	return &AuthHandler{
		authService,
	}
}

func (ah *AuthHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request entity.RegisterRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		_, err := ah.authService.Register(request)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Register Sukses",
			"Status":  http.StatusOK,
		})
	}
}

func (ah *AuthHandler) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request entity.LoginRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		login, err := ah.authService.Login(request)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Berhasil login",
			"Data":    login,
			"Status":  http.StatusOK,
		})
	}
}
