package main

import (
	"booking-room/initializers"
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
	
}


