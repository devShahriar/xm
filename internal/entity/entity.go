package entity

import (
	"time"

	"github.com/google/uuid"
)

// Company represents the core domain model with Gorm annotations
type Company struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name         string    `gorm:"type:varchar(15);unique" json:"name"`
	Description  string    `gorm:"type:varchar(3000)" json:"description"`
	NumEmployees int       `gorm:"type:int" json:"num_employees"`
	Registered   bool      `gorm:"type:boolean" json:"registered"`
	Type         string    `gorm:"type:varchar(50)" json:"type"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
