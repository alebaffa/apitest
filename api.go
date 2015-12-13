package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", GetUserBasicInfo)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// GetUserBasicInfo retrieves the most used information about the user given in input
func GetUserBasicInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ciao")
}
