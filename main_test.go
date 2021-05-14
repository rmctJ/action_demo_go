package main

import (
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/pi", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}

	expected := fmt.Sprint(math.Pi)
	if rr.Body.String() != expected {
		t.Errorf(
			"unexpected:got (%v) want (%v)",
			rr.Body.String(),
			expected,
		)
	}
	

}