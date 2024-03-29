package handler

import "github.com/labstack/echo/v4"

func (ah *AuthHandler) RegisterRoutes(e *echo.Echo) {
	//adminGroup := e.Group("/admin")
	//adminGroup.POST("/login", ah.LoginAdmin())

	userGroup := e.Group("/user")
	userGroup.POST("/register", ah.Register())
	userGroup.POST("/login", ah.LoginUser())
}
