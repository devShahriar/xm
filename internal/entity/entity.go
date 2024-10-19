package entity

import (
	"time"

	"github.com/google/uuid"
)

// Company represents the core domain model with Gorm annotations
type Company struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	Name         string    `gorm:"type:varchar(15);unique"`
	Description  string    `gorm:"type:varchar(3000)"`
	NumEmployees int       `gorm:"type:int"`
	Registered   bool      `gorm:"type:boolean"`
	Type         string    `gorm:"type:varchar(50)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
