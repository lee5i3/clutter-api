package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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
	fmt.Println("Listening....")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	handleRequests()
}
