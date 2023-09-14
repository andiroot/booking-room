package room

import (
	"booking-room/room"
	"gorm.io/gorm"
)

type Repository interface {
	FindAllRoom(UserID uint) ([]Room, error)
	FindAllByUser(UserID uint) ([]Room, error)
	FindByID(ID int) (Room, error)
	Create(room Room) (Room, error)
	Update(room Room) (Room, error)
	Delete(room Room) (Room, error)
}

type repository struct {
	db *gorm.DB
}

// FindAllRoom implements Repository.
func (*repository) FindAllRoom(UserID uint) ([]room.Room, error) {
	panic("unimplemented")
}

// Update implements Repository.
func (*repository) Update(room room.Room) (room.Room, error) {
	panic("unimplemented")
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db, // Assign the provided db to the struct field
	}
}

func (r *repository) FindAll(UserID uint) ([]Room, error) {
	var rooms []Room
	err := r.db.Find(&rooms).Error

	return rooms, err
}

func (r *repository) FindByID(ID int) (Room, error) {
	var room Room
	err := r.db.First(&room, ID).Error
	return room, err
}

func (r *repository) Create(room Room) (Room, error) {
	err := r.db.Create(&room).Error
	return room, err
}

func (r *repository) Delete(room Room) (Room, error) {
	err := r.db.Delete(&room).Error
	return room, err
}

func (r *repository) FindAllByUser(UserID uint) ([]Room, error) {
	var rooms []Room
	err := r.db.Where("user_id = ?", UserID).Find(&rooms).Error
	return rooms, err
}
