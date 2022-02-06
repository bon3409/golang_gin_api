package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB
var err error

var user = os.Getenv("DATABASE_USER")
var host = os.Getenv("DATABASE_HOST")
var password = os.Getenv("DATABASE_PASSWORD")
var port = os.Getenv("DATABASE_PORT")
var database = os.Getenv("DATABASE_NAME")
var connection string

func DbConfiguration() {
	connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
}
