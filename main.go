package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	value := os.Getenv("VAR")

	fmt.Fprintf(w, fmt.Sprintf("Welcome to the %s!", value))
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func main() {
	handleRequests()
}
