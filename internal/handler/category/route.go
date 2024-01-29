package category

import "github.com/labstack/echo/v4"

func (ah *CategoryHandler) RegisterRoutes(e *echo.Echo) {

	categoryGroup := e.Group("/category")
	categoryGroup.POST("/add", ah.CreateCategory)
	categoryGroup.GET("/get", ah.GetCategoryByID)
	categoryGroup.PUT("/update", ah.UpdateCategory)
	categoryGroup.DELETE("/delete", ah.DeleteCategory)
	categoryGroup.GET("/getall", ah.GetAllCategories)
}
