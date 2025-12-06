package dto

type MessageInput struct {
	Content string `json:"content" binding:"required"`
	RoomID  uint   `json:"room_id" binding:"required"`
}

type MessageGetInput struct {
	RoomID uint `json:"room_id" binding:"required"`
}

type MessageOutput struct {
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}
