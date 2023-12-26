package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {
	var err error
	host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbname)

	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(dsn)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database !")
	}

}
