package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemInfo(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Fatal(err)
	}
	NewRouter().ServeHTTP(w, req)
	assert.Equal(t, "ciao", w.Body.String(), "they should be equal")
}

func TestGetUserInfo(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user/123", nil)
	if err != nil {
		log.Fatal(err)
	}
	NewRouter().ServeHTTP(w, req)
	assert.Equal(t, "User 123", w.Body.String(), "they should be equal")
}
