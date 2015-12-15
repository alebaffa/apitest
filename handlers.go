package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetSystemInfo retrieves the information about the system
func GetSystemInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ciao")
}

func GetRandomUser(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Name: "Alessandro Baffa"},
		User{Name: "A random user"},
	}
	json.NewEncoder(w).Encode(users)
}

// GetUserInfo retrieves the information about a specific user
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	fmt.Println("utente inviato è ", userId)
	fmt.Fprintf(w, "User %s", userId)
}
