package models

import (
	"time"
)

// Log represents a log in the chat platform.
type Log struct {
	ID         uint `gorm:"primaryKey"`
	APIName    string
	APIUrl     string
	APIMethod  string
	APISuccess bool
	UserID     string
	IP         string
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

// LogError represents a log error in the chat platform.
type LogError struct {
	ID        uint `gorm:"primaryKey"`
	ErrorCode string
	LogID     int
	Msg       string
}
