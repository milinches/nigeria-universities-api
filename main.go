package main

import (
	"log"
	"net/http"
	"github.com/milinches/nigeria-universities-api/controllers"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.GetUniversities).Methods("GET")
	router.HandleFunc("/{abbreviation}", controllers.GetSpecificUniversity).Methods("GET")
	log.Println("Starting server at :7070")
	err := http.ListenAndServe(":7070", router)
	log.Fatal(err)
}