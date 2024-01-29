package category

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"technopartner/internal/entity"
	"time"
)

// CategoryRepo defines the interface for category repository.
type CategoryRepo interface {
	CreateCategory(name, description string) (*entity.Category, error)
	GetCategoryByID(categoryID uuid.UUID) (*entity.Category, error)
	UpdateCategory(categoryID uuid.UUID, name, description string) (*entity.Category, error)
	DeleteCategory(categoryID uuid.UUID) error
	GetAllCategories() ([]entity.Category, error)
}

// categoryRepo is the GORM implementation of the CategoryRepo interface.
type categoryRepo struct {
	db *gorm.DB
}

// NewCategoryRepo creates a new instance of CategoryRepo.
func NewCategoryRepo(db *gorm.DB) CategoryRepo {
	return &categoryRepo{db}
}

// CreateCategory adds a new category to the database.
func (r *categoryRepo) CreateCategory(name, description string) (*entity.Category, error) {
	category := &entity.Category{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := r.db.Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

// GetCategoryByID retrieves a category by its ID.
func (r *categoryRepo) GetCategoryByID(categoryID uuid.UUID) (*entity.Category, error) {
	category := &entity.Category{}
	if err := r.db.First(category, "id = ?", categoryID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return category, nil
}

// UpdateCategory updates the details of an existing category.
func (r *categoryRepo) UpdateCategory(categoryID uuid.UUID, name, description string) (*entity.Category, error) {
	category, err := r.GetCategoryByID(categoryID)
	if err != nil {
		return nil, err
	}

	category.Name = name
	category.Description = description
	category.UpdatedAt = time.Now()

	if err := r.db.Save(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

// DeleteCategory deletes a category by its ID.
func (r *categoryRepo) DeleteCategory(categoryID uuid.UUID) error {
	return r.db.Delete(&entity.Category{}, "id = ?", categoryID).Error
}

// GetAllCategories retrieves all categories from the database.
func (r *categoryRepo) GetAllCategories() ([]entity.Category, error) {
	var categories []entity.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
