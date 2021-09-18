package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/milinches/nigeria-universities-api/models"
)

func GetUniversities(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.Write([]byte("<div><h1>4$4 Not Found, Boss</h1></div>"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Uni())
}

func GetSpecificUniversity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, getItem := range models.Uni() {
		if getItem.Abbrv == params["abbreviation"] {
			json.NewEncoder(w).Encode(getItem)
			return
		}
	}
	w.Write([]byte("Sorry, this is not Available"))
}
