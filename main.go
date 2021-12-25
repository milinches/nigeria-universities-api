package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/milinches/nigeria-universities-api/routes"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	port, exist := os.LookupEnv("PORT")
	if !exist {
		log.Fatal("No .env file found")
	}
	router := routes.NewRouter()
	log.Println("Starting server at :" + port)
	err := http.ListenAndServe(":"+port, router)
	log.Fatal(err)
}
