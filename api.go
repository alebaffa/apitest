package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getUserBasicInfo)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func getUserBasicInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ciao")
}
