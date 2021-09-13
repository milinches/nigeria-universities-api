package main

import (
	"log"
	"net/http"
	"github.com/milinches/nigeria-universities-api/routes"
)

func main(){
	router := routes.NewRouter()
	log.Println("Starting server at :7070")
	err := http.ListenAndServe(":7070", router) 	
	log.Fatal(err)
}