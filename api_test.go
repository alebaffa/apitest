package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	GetUserBasicInfo(w, req)
	assert.Equal(t, "ciao", w.Body.String(), "they should be equal")
}
