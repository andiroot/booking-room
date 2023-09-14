package handler

import (
	"booking-room/room"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type roomHandler struct {
	roomService room.Service
}

func NewRoomHandler(service room.Service) *roomHandler {
	return &roomHandler{service}
}
func (h *roomHandler) PostRoomsHandler(c *gin.Context) {
	var roomRequest room.RoomRequest
	err := c.ShouldBindJSON(&roomRequest)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s,condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}
	jwtClaims, _ := c.Get("jwtClaims")
	claims, _ := jwtClaims.(jwt.MapClaims)
	userID, _ := claims["sub"].(float64)
	room, err := h.roomService.Create(roomRequest, uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": room,
	})
}

func (h *roomHandler) GetRooms(c *gin.Context) {
	jwtClaims, _ := c.Get("jwtClaims")
	claims, _ := jwtClaims.(jwt.MapClaims)
	userID, _ := claims["sub"].(uint)
	rooms, err := h.roomService.FindAllRoom(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	var roomsResponse []room.RoomResponse
	for _, b := range rooms {
		roomResponse := room.ConvertToRoomResponse(b)
		roomsResponse = append(roomsResponse, roomResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": roomsResponse,
	})
}

func (h *roomHandler) UpdateRoomHandler(c *gin.Context) {
	var roomRequest room.RoomRequest
	err := c.ShouldBindJSON(&roomRequest)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s,condition: %s", e.Field(), e.ActualTag())

				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}
	ID, _ := strconv.Atoi(c.Param("id"))
	b, err := h.roomService.Update(ID, roomRequest)
	roomResponse := room.ConvertToRoomResponse(b)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": roomResponse,
	})
}

func (h *roomHandler) GetRoom(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	b, err := h.roomService.FindByID(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err})
		return
	}
	roomResponse := room.ConvertToRoomResponse(b)
	c.JSON(http.StatusBadRequest, gin.H{
		"data": roomResponse,
	})
}

func (h *roomHandler) DeleteRoom(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	b, err := h.roomService.Delete(ID)
	roomResponse := room.ConvertToRoomResponse(b)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": roomResponse,
	})
}
