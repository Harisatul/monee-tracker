package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"technopartner/db"
	"technopartner/internal/handler"
	h "technopartner/internal/handler/category"
	"technopartner/internal/repository/category"
	"technopartner/internal/repository/user"
	cs "technopartner/internal/service/category"
	user2 "technopartner/internal/service/user"
)

func main() {
	db.Init()

	d := db.DB
	authRepo := user.New(d)
	authUsecase := user2.New(authRepo)
	authHandler := handler.New(authUsecase)

	categoryRepo := category.NewCategoryRepo(d)
	categoryService := cs.NewCategoryService(categoryRepo)
	categoryHandler := h.NewCategoryHandler(categoryService)

	e := echo.New()
	//log.LogMiddleware(e)
	//e.Use(middleware.CORS())

	authHandler.RegisterRoutes(e)
	categoryHandler.RegisterRoutes(e)
	err := e.Start(":9999")

	if err != nil {
		fmt.Print(err.Error())
	}
}
