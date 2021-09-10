package main

import (
	"log"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Api will be rendered here"))
}

func getByAbbrv(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Get school by abbrv"))
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", api)
	mux.HandleFunc("/abbrv", getByAbbrv)
	log.Println("Starting server at :7070")
	err := http.ListenAndServe(":7070", mux)
	log.Fatal(err)
}