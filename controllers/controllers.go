package controllers

import (
	"MileTravel/database"
	"MileTravel/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func Testimonials(w http.ResponseWriter, r *http.Request) {
	var testimonials []models.Testimonial
	database.DB.Find(&testimonials)

	json.NewEncoder(w).Encode(testimonials)

}

func getVarFromPath(r *http.Request, varName string) string {
	vars := mux.Vars(r)
	value := vars[varName]
	return value
}

func TestimonialById(w http.ResponseWriter, r *http.Request) {
	id := getVarFromPath(r, "id")
	var testimonial models.Testimonial
	database.DB.First(&testimonial, id)
	json.NewEncoder(w).Encode(testimonial)
}

func CreateTestimonial(w http.ResponseWriter, r *http.Request) {
	var testimonial models.Testimonial
	json.NewDecoder(r.Body).Decode(&testimonial)
	database.DB.Create(&testimonial)
	json.NewEncoder(w).Encode(testimonial)
}

func DeleteTestimonial(w http.ResponseWriter, r *http.Request) {
	id := getVarFromPath(r, "id")
	var testimonial models.Testimonial
	database.DB.Delete(&testimonial, id)
	json.NewEncoder(w).Encode(testimonial)
}

func UpdateTestimonial(w http.ResponseWriter, r *http.Request) {
	id := getVarFromPath(r, "id")
	var testimonial models.Testimonial
	database.DB.First(&testimonial, id)
	json.NewDecoder(r.Body).Decode(&testimonial)
	database.DB.Save(&testimonial)
	json.NewEncoder(w).Encode(testimonial)
}
