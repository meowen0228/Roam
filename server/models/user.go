package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the chat platform.
type User struct {
	ID        uint           `gorm:"primaryKey"`
	Username  string         `gorm:"unique;not null"`
	Email     string         `gorm:"unique;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
