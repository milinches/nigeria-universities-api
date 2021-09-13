package main

import (
	"log"
	"net/http"
	"github.com/milinches/nigeria-universities-api/controllers"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.GetUniversities)
	mux.HandleFunc("/controllers/abbreviation", controllers.GetSpecificUniversity)
	log.Println("Starting server at :7070")
	err := http.ListenAndServe(":7070", mux)
	log.Fatal(err)
}