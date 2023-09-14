package initializers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error){
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s",
        "root",
        "password",
        "127.0.0.1:3306",
        "pustakaapi",
        "utfmb4",
        "True",
        "Local",
    )
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}