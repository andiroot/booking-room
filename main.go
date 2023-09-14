package main

import (
	"booking-room/handler"
	"booking-room/initializers"
	"booking-room/middleware"
	"booking-room/room"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	db, err := initializers.ConnectToDatabase()
	if err != nil {
		panic("connect db error euy")
	}

	initializers.SyncDatabase(db)
}

func main() {
	db, err := initializers.ConnectToDatabase()
	if err != nil {
		panic("connect db error euy")
	}
	router := gin.Default()
	routerV1 := router.Group("/v1")
	routerV1Rooms := routerV1.Group("/rooms", middleware.RequireAuth)
	
	roomRepository := room.NewRepository(db)
	roomService := room.NewService(roomRepository)
	roomHandler := handler.NewRoomHandler(roomService)

	routerV1Rooms.POST("", roomHandler.PostRoomsHandler)
	routerV1Rooms.GET("", roomHandler.GetRooms)
	routerV1Rooms.GET("/:id", roomHandler.GetRoom)
	routerV1Rooms.PUT("/:id", roomHandler.UpdateRoomHandler)
	routerV1Rooms.DELETE("/:id", roomHandler.DeleteRoom)
}


