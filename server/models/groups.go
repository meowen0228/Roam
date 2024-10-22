package models

import (
	"time"

	"gorm.io/gorm"
)

// Group represents a group in the chat platform.
type Group struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// GroupMembers 群組成員
type GroupMembers struct {
	ID         uint `gorm:"primaryKey"`
	GroupRefer int
	Group      Group `gorm:"foreignKey:GroupRefer"`
	UserRefer  int
	User       User           `gorm:"foreignKey:UserRefer"`
	JoinedAt   time.Time      `gorm:"autoCreateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// GroupMessages 群組訊息
type GroupMessages struct {
	ID         uint `gorm:"primaryKey"`
	GroupRefer int
	Group      Group `gorm:"foreignKey:GroupRefer"`
	SenderID   int
	User       User           `gorm:"foreignKey:SenderID"`
	Content    string         `gorm:"type:text;not null"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
