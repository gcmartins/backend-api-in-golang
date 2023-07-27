package main

import (
	"MileTravel/database"
	"MileTravel/routes"
	"fmt"
)

func main() {
	database.DbConfig()

	fmt.Println("Starting server...")
	router := routes.LoadRouter()
	router.Run()
}
