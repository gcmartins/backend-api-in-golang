package controllers_test

import (
	"MileTravel/database"
	"MileTravel/models"
	"MileTravel/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	router := routes.LoadRouter()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	expected_output := `{"message":"Home"}`

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, expected_output, res.Body.String())
}

func loadDatabase() {
	testimonials := []models.Testimonial{
		{User: "User1", Image: "Image1", Description: "Description1"},
		{User: "User2", Image: "Image2", Description: "Description2"},
	}

	for _, testimonial := range testimonials {
		database.DB.Create(&testimonial)
	}
}

func TestTestimonials(t *testing.T) {
	database.TestDbConfig()
	loadDatabase()
	defer database.ClearTestDb()
	router := routes.LoadRouter()

	t.Run("Create Testimonial", func(t *testing.T) {
		expectedTestimonial := models.Testimonial{User: "John", Image: "Image 1", Description: "Description 1"}
		testimonialJson, _ := json.Marshal(expectedTestimonial)
		req, _ := http.NewRequest(http.MethodPost, "/api/testimonials", bytes.NewBuffer(testimonialJson))

		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		var createdTestimonial models.Testimonial
		json.Unmarshal(res.Body.Bytes(), &createdTestimonial)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expectedTestimonial.User, createdTestimonial.User)
		assert.Equal(t, expectedTestimonial.Image, createdTestimonial.Image)
		assert.Equal(t, expectedTestimonial.Description, createdTestimonial.Description)
	})

	t.Run("Retrieve Testimonial", func(t *testing.T) {
		ids := []string{"1", "2", "100"}

		for _, id := range ids {
			req, _ := http.NewRequest(http.MethodGet, "/api/testimonials/"+id, nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			var retrievedTestimonial models.Testimonial

			if id == "100" {
				assert.Equal(t, http.StatusNotFound, res.Code)
				assert.Equal(t, `{"message":"Resource not found."}`, res.Body.String())
			} else {
				json.Unmarshal(res.Body.Bytes(), &retrievedTestimonial)
				assert.Equal(t, http.StatusOK, res.Code)
				assert.Equal(t, "User"+id, retrievedTestimonial.User)
				assert.Equal(t, "Image"+id, retrievedTestimonial.Image)
				assert.Equal(t, "Description"+id, retrievedTestimonial.Description)
			}
		}
	})

	t.Run("Delete Testimonial", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/api/testimonials/3", nil)

		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, `{"message":"Resource deleted."}`, res.Body.String())
	})

	t.Run("Update Testimonial", func(t *testing.T) {
		testimonialJson := `{"user": "Mary", "image": "Image 30"}`
		req, _ := http.NewRequest(http.MethodPut, "/api/testimonials/1", bytes.NewBufferString(testimonialJson))

		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		var updatedTestimonial models.Testimonial
		json.Unmarshal(res.Body.Bytes(), &updatedTestimonial)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, 1, updatedTestimonial.Id)
		assert.Equal(t, "Mary", updatedTestimonial.User)
		assert.Equal(t, "Image 30", updatedTestimonial.Image)
		assert.Equal(t, "Description1", updatedTestimonial.Description)
	})

	t.Run("Update inexistent Testimonial", func(t *testing.T) {
		testimonialJson := `{"user": "Mary", "image": "Image 30"}`
		req, _ := http.NewRequest(http.MethodPut, "/api/testimonials/20", bytes.NewBufferString(testimonialJson))

		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, `{"message":"Resource not found."}`, res.Body.String())
	})

}
