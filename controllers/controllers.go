package controllers

import (
	"MileTravel/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func Testimonials(w http.ResponseWriter, r *http.Request) {
	testimonials := []models.Testimony{
		{Id: 1, User: "Bob", Image: "image1", Description: "decription1"},
		{Id: 2, User: "Jane", Image: "image2", Description: "decription2"},
	}

	json.NewEncoder(w).Encode(testimonials)

}
