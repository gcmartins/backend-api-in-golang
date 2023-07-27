package routes_test

import (
	"MileTravel/routes"
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
