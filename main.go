package main

import (
	"MileTravel/routes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
