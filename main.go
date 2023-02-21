package main

import (
	"apiBlog/database"
	"apiBlog/model"
	"apiBlog/routes"
	"apiBlog/seed"
	"log"
)

func main() {
	database.InitConnectDB()
	model.Delete()
	model.Migration()
	seed.Seed()
	routes.UserRoutes()
	log.Println("hello word")
}

