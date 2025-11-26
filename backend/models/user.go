package models

type User struct {
	Username string `gorm:"uniqueIndex;not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Rooms    []Room `gorm:"many2many:user_rooms;"`
}
