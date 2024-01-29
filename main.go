package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"technopartner/db"
	"technopartner/internal/handler"
	"technopartner/internal/repository"
	"technopartner/internal/service"
)

func main() {
	db.Init()

	authRepo := repository.New(db.DB)
	authUsecase := service.New(authRepo)
	authHandler := handler.New(authUsecase)

	e := echo.New()
	//log.LogMiddleware(e)
	//e.Use(middleware.CORS())

	authHandler.RegisterRoutes(e)

	err := e.Start(":9999")

	if err != nil {
		fmt.Print(err.Error())
	}
}
