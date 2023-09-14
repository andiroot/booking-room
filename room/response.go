package room

import "time"

type RoomResponse struct {
	// ID          int    `json:"id"`
	// Title       string `json:"title"`
	// Price       int    `json:"price"`
	// Description string `json:"description"`
	// Rating      int    `json:"rating"`
	ID int `json:"id"`
	Name string `json:"name"`
	Capacity int `json:"capacity"`
	Status bool `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}