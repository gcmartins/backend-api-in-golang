package routes

import (
	"MileTravel/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/testimonials", controllers.Testimonials)
}
