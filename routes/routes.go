package routes

import (
	"MileTravel/controllers"

	"github.com/gorilla/mux"
)

func LoadRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Index)
	r.HandleFunc("/api/testimonials", controllers.Testimonials).Methods("Get")
	r.HandleFunc("/api/testimonials", controllers.CreateTestimonial).Methods("Post")
	r.HandleFunc("/api/testimonials/{id}", controllers.TestimonialById).Methods("Get")
	r.HandleFunc("/api/testimonials/{id}", controllers.DeleteTestimonial).Methods("Delete")
	r.HandleFunc("/api/testimonials/{id}", controllers.UpdateTestimonial).Methods("Put")

	return r
}
