package category

import (
	"github.com/google/uuid"
	"technopartner/internal/entity"
	"technopartner/internal/repository/category"
)

// CategoryService defines the interface for category service.
type CategoryService interface {
	CreateCategory(name, description string) (*entity.Category, error)
	GetCategoryByID(categoryID uuid.UUID) (*entity.Category, error)
	UpdateCategory(categoryID uuid.UUID, name, description string) (*entity.Category, error)
	DeleteCategory(categoryID uuid.UUID) error
	GetAllCategories() ([]entity.Category, error)
}
type categoryService struct {
	categoryRepo category.CategoryRepo
}

func NewCategoryService(categoryRepo category.CategoryRepo) *categoryService {
	return &categoryService{
		categoryRepo,
	}
}

// CreateCategory adds a new category using the repository.
func (s *categoryService) CreateCategory(name, description string) (*entity.Category, error) {
	return s.categoryRepo.CreateCategory(name, description)
}

// GetCategoryByID retrieves a category by its ID using the repository.
func (s *categoryService) GetCategoryByID(categoryID uuid.UUID) (*entity.Category, error) {
	return s.categoryRepo.GetCategoryByID(categoryID)
}

// UpdateCategory updates the details of an existing category using the repository.
func (s *categoryService) UpdateCategory(categoryID uuid.UUID, name, description string) (*entity.Category, error) {
	return s.categoryRepo.UpdateCategory(categoryID, name, description)
}

// DeleteCategory deletes a category by its ID using the repository.
func (s *categoryService) DeleteCategory(categoryID uuid.UUID) error {
	return s.categoryRepo.DeleteCategory(categoryID)
}

// GetAllCategories retrieves all categories from the repository.
func (s *categoryService) GetAllCategories() ([]entity.Category, error) {
	return s.categoryRepo.GetAllCategories()
}
