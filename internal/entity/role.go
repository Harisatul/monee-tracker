package entity

import (
	"time"
)

type Role struct {
	ID        string `gorm:"type:text;primaryKey"`
	Role      string `json:"Role" form:"Role"`
	Users     User
	CreatedAt time.Time
	UpdatedAt time.Time
}
