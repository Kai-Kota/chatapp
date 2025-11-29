package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Member1 uint `gorm:"not null"`
	Member2 uint `gorm:"not null"`
	//Messages []Message `gorm:"many2many:message_rooms"`
}
