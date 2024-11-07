package models

import (
	"time"

	"gorm.io/gorm"
)

// PrivateMessage represents a private message between users.
type PrivateMessage struct {
	ID         uint `gorm:"primaryKey"`
	SenderID   int
	SenderUser User `gorm:"foreignKey:SenderID"`

	ReceiverID   int
	ReceiverUser int            `gorm:"foreignKey:ReceiverID"`
	Content      string         `gorm:"type:text;not null"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	IsRead       bool           `gorm:"default:false"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
