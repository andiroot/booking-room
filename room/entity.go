package room

import "time"

type Room struct{
	ID int
	Name string
	Capacity int
	Status bool
	CreatedAt time.Time
	UpdatedAt time.Time
}