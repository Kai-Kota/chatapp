package dto

type CreateRoomInput struct {
	RoomID int `json:"room_id" binding:"required"`
}
