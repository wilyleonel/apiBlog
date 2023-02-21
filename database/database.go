package database

import (
	"fmt"
	"log"

	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DNS = "host=localhost user=postgres password=admin dbname=blog port=5432"

func InitConnectDB() {
	var error error
	DB, error = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		fmt.Println("http://localhost:3000")
		log.Println("Database initialized")
	}
}

// type Server struct {
// 	DB *gorm.DB
// 	// Router *mux.Router
// }

// func (server *Server) initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
// 	var err error
// 	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
// 	server.DB, err = gorm.Open(DBURL)
// 	if err != nil {
// 		fmt.Printf("Cannot connect to %s database", Dbdriver)
// 		log.Fatal("This is the error:", err)
// 	} else {
// 		fmt.Printf("We are connected to the %s database", Dbdriver)
// 	}
// }

// func Run() {
// 	var err error
// 	err = godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error getting env, %v", err)
// 	} else {
// 		fmt.Println("We are getting the env values")
// 	}
// 	server.initialize
// }
