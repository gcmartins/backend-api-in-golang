package routes_test

import (
	"MileTravel/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexRoute(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Index)

	handler.ServeHTTP(res, req)

	expected_output := "Home"

	if res.Code != http.StatusOK {
		t.Error("expected", http.StatusOK, "got", res.Code)
	}

	if res.Body.String() != expected_output {
		t.Error("expected", expected_output, "got", res.Body.String())
	}

}
