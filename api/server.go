package api

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/milinches/nigeria-universities-api/api/router"
)

func Run() {
	envErr := godotenv.Load(".file.env")
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	log.Printf("\nClick here to open server => http://localhost:%s/api/v1", port)
	r := router.NEW()
	log.Fatal(http.ListenAndServe(":"+port, r))
}
