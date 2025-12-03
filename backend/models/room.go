package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name     string    `gorm:"not null"`
	Messages []Message `gorm:"many2many:room_messages"`
}
	