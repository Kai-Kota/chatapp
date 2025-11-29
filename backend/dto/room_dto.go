package dto

type CreateRoomInput struct {
	Pertner string `json:"pertner" binding:"required"`
}
