package category

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"technopartner/internal/service/category"
)

type CategoryHandler struct {
	categoryService category.CategoryService
}

func NewCategoryHandler(authService category.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		authService,
	}
}

// CreateCategory handles the HTTP request to create a new category.
func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	category, err := h.categoryService.CreateCategory(input.Name, input.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create category")
	}

	return c.JSON(http.StatusCreated, category)
}

// GetCategoryByID handles the HTTP request to get a category by its ID.
func (h *CategoryHandler) GetCategoryByID(c echo.Context) error {
	categoryID, err := uuid.Parse(c.QueryParam("categoryId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}

	category, err := h.categoryService.GetCategoryByID(categoryID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Category not found")
	}

	return c.JSON(http.StatusOK, category)
}

// UpdateCategory handles the HTTP request to update a category.
func (h *CategoryHandler) UpdateCategory(c echo.Context) error {
	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	categoryID, err := uuid.Parse(c.QueryParam("categoryId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}

	category, err := h.categoryService.UpdateCategory(categoryID, input.Name, input.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update category")
	}

	return c.JSON(http.StatusOK, category)
}

// DeleteCategory handles the HTTP request to delete a category.
func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	categoryID, err := uuid.Parse(c.QueryParam("categoryId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}

	err = h.categoryService.DeleteCategory(categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete category")
	}

	return c.NoContent(http.StatusOK)
}

func (h *CategoryHandler) GetAllCategories(c echo.Context) error {
	categories, err := h.categoryService.GetAllCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get categories")
	}

	return c.JSON(http.StatusOK, categories)
}
