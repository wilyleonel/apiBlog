package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}
var DB *gorm.DB

func InitConnectDB() {
	err:=godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err.Error())
		return
	}
	connInfo := connection{
		Host:     os.Getenv("POSTGRES_URL"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}
	var error error
	DB, error = gorm.Open(postgres.Open(connToString(connInfo)), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		fmt.Println("http://localhost:3000")
		log.Println("Database initialized")
	}
}

func connToString(info connection) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)

}