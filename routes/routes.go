package routes

import (
	"MileTravel/controllers"

	"github.com/gorilla/mux"
)

func LoadRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Index)
	r.HandleFunc("/api/testimonials", controllers.Testimonials)

	return r
}
