package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	RoomID   int `gorm:"uniqueIndex;not null"`
	Messages []Message
	Users    []User
}
