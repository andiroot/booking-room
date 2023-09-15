package main

import (
	"booking-room/handler"
	"booking-room/initializers"
	"booking-room/middleware"
	"booking-room/room"
	"booking-room/schedule"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

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

	roomRepository := room.NewRepository(db)
	roomService := room.NewService(roomRepository)
	roomHandler := handler.NewRoomHandler(roomService)
	scheduleRepository := schedule.NewRepository(db)
	scheduleService := schedule.NewService(scheduleRepository)
	scheduleHandler := handler.NewScheduleHandler(scheduleService)

	router := gin.Default()
	routerV1 := router.Group("/v1")
	routerV1Rooms := routerV1.Group("/rooms", middleware.RequireAuth)
	routerV1Rooms.POST("", roomHandler.PostRoomsHandler)
	routerV1Rooms.GET("", roomHandler.GetRooms)
	routerV1Rooms.GET("/:id", roomHandler.GetRoom)
	routerV1Rooms.PUT("/:id", roomHandler.UpdateRoomHandler)
	routerV1Rooms.DELETE("/:id", roomHandler.DeleteRoom)

	routerV1Books := routerV1.Group("/schedule", middleware.RequireAuth)
	routerV1Books.POST("", scheduleHandler.PostScheduleHandler)
	routerV1Books.GET("", scheduleHandler.GetSchedules)
	routerV1Books.GET("/:id", scheduleHandler.GetSchedule)
	routerV1Books.PUT("/:id", scheduleHandler.UpdateScheduleHandler)
	routerV1Books.DELETE("/:id", scheduleHandler.DeleteSchedule)

	routerUser := router.Group("/v1/user")
	routerUser.POST("/signup", userHandler.Signup)
	routerUser.POST("/login", userHandler.Login)
}
