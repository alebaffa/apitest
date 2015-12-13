package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", GetSystemInfo)
	router.HandleFunc("/user/{userId}", GetUserInfo)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

// GetSystemInfo retrieves the information about the system
func GetSystemInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ciao")
}

// GetUserInfo retrieves the information about a specific user
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	fmt.Println("utente inviato è ", userId)
	fmt.Fprintf(w, "User %s", userId)
}
