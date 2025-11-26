package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	MessageID int    `gorm:"uniqueIndex;not null"`
	Content   string `gorm:"not null"`
	UserID    int    `gorm:"not null"`
}