package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Subcategory represents the subcategory entity.
type Subcategory struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"subcategoryId"`
	Name        string    `json:"subcategoryName"`
	Description string    `json:"subcategoryDescription"`
	CategoryID  uuid.UUID `gorm:"type:uuid;index" json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (subcategory *Subcategory) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	subcategory.ID = uuid.New()
	return
}
