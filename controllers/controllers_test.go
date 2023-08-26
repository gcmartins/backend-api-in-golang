package controllers_test

import (
	"MileTravel/database"
	"MileTravel/models"
	"MileTravel/routes"
	"MileTravel/storage"
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func loadDatabase(rows int) {
	for i := 1; i <= rows; i++ {
		testimonial := models.Testimonial{
			User:        "User" + strconv.Itoa(i),
			Image:       "Image" + strconv.Itoa(i),
			Description: "Description" + strconv.Itoa(i),
		}
		database.DB.Create(&testimonial)
	}
}

func TestTestimonials(t *testing.T) {
	database.TestDbConfig()
	loadDatabase(10)
	defer database.ClearTestDb()
	router := routes.LoadRouter()
	storage.SetupTestStorage()
	defer storage.ClearTestStorage()

	t.Run("Retrieve All Testimonials", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/testimonials", nil)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var testimonials []models.Testimonial

		json.Unmarshal(res.Body.Bytes(), &testimonials)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, 10, len(testimonials))
	})

	t.Run("Create Testimonial", func(t *testing.T) {
		expectedTestimonial := models.Testimonial{User: "John", Image: "Image 1", Description: "Description 1"}
		testimonialJson, _ := json.Marshal(expectedTestimonial)

		imageFile, _ := os.Open("../foto.jpg")
		defer imageFile.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		jsonPart, _ := writer.CreateFormField("json")
		jsonPart.Write(testimonialJson)

		imagePart, _ := writer.CreateFormFile("image", "image.jpg")
		io.Copy(imagePart, imageFile)

		writer.Close()

		req, _ := http.NewRequest(http.MethodPost, "/api/testimonials", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

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

	t.Run("Retrieve Testimonial Home", func(t *testing.T) {
		parseTestimonials := func(res *httptest.ResponseRecorder, req *http.Request) []models.Testimonial {
			router.ServeHTTP(res, req)
			var testimonials []models.Testimonial
			json.Unmarshal(res.Body.Bytes(), &testimonials)
			return testimonials
		}

		req, _ := http.NewRequest(http.MethodGet, "/api/testimonials-home", nil)
		res := httptest.NewRecorder()
		testimonials := parseTestimonials(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, 3, len(testimonials))

		res2 := httptest.NewRecorder()
		testimonials2 := parseTestimonials(res2, req)
		assert.Equal(t, http.StatusOK, res2.Code)
		assert.Equal(t, 3, len(testimonials2))

		assert.NotEqualValues(t, testimonials, testimonials2)
	})

}
