package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Category represents the category entity.
type Category struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"categoryId"`
	Name        string    `json:"categoryName"`
	Description string    `json:"categoryDescription"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	category.ID = uuid.New()
	return
}
