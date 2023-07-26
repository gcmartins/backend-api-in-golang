package main

import (
	"MileTravel/database"
	"MileTravel/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	database.DbConfig()

	fmt.Println("Starting server...")
	router := routes.LoadRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
