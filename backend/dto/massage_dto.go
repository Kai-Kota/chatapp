package dto

type messageInput struct {
	RoomID  uint   `json:"room_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}