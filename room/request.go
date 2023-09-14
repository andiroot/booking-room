package room

type RoomRequest struct {
	Name string `binding:"required"`
	Capacity int `json:"price" validate:"required,number"`
	Status bool `binding:"required"`
}