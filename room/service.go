package room

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service interface {
	FindAll(UserID uint) ([]Room, error)
	FindByID(ID int) (Room, error)
	Create(roomRequest RoomRequest, UserID uint) (Room, error)
	Update(ID int, roomRequest RoomRequest) (Room, error)
	Delete(ID int) (Room, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) FindAll() ([]Room, error) {
	rooms, err := s.repository.FindAll()
	return rooms, err
}
func (s *service) FindByID(ID int) (Room, error) {
	room, err := s.repository.FindByID(ID)
	return room, err
}
func (s *service) Create(roomRequest RoomRequest, UserID uint) (Room, error) {
	room := Room{
		Name:       roomRequest.Name,
		Capacity:       roomRequest.Capacity,
		Status: roomRequest.Status,
	}
	newroom, err := s.repository.Create(room)
	return newroom, err
}
func (s *service) Update(ID int, roomRequest RoomRequest) (Room,
	error) {
	room, err := s.repository.FindByID(ID)
	if roomRequest.Name != "" {
		room.Name = roomRequest.Name
	}
	if roomRequest.Capacity != 0 {
		room.Capacity = roomRequest.Capacity
	}
	if roomRequest.Status != false {
		room.Status =
			roomRequest.Status
	}
	
	newroom, err := s.repository.Update(room)
	return newroom, err
}

func (s *service) Delete(ID int) (Room, error) {
	room, err := s.repository.FindByID(ID)
	_, err = s.repository.Delete(room)
	return room, err
}

func (h *roomHandler) Getrooms(c *gin.Context) {
	rooms, err := h.roomService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": rooms,
	})
}