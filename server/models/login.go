package models

import (
	"time"
)

// LoginInfo represents a login record.
type LoginInfo struct {
	ID        uint `gorm:"primaryKey"`
	UserID    int
	User      User      `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
