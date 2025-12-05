package services

type IMessageService interface {
	CreateMessage(roomID uint, content string)
}
