package main

import (
	"MileTravel/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	router := routes.LoadRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
