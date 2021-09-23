package main

import (
	"log"
	"net/http"

	"github.com/milinches/nigeria-universities-api/routes"
)

func main() {
	router := routes.NewRouter()
	log.Println("Starting server at :8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
