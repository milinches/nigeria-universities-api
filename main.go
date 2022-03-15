package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/milinches/nigeria-universities-api/routes"
)

func main() {
	envErr := godotenv.Load(".file.env")
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	router := routes.NewRouter()

	log.Println("Starting server...")
	err := http.ListenAndServe(":"+port, router)
	log.Fatal(err)
}
