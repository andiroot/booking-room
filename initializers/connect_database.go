package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error){
    dbUsername := os.Getenv("DB_USERNAME")
    dbPass := os.Getenv("DB_PASS")
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s",
        dbUsername,
        dbPass,
        dbHost,
        dbName,
        "utf8mb4",
        "True",
        "Local",
    )
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}