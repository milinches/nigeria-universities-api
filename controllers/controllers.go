package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/milinches/nigeria-universities-api/models"
)

func GetUniversities(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"{
		w.Write([]byte("<div><h1>4$4 Not Found, Boss</h1></div>"))
		return
	}
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(401)
		w.Write([]byte("Method not Allowed"))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Uni())
}

func GetSpecificUniversity(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	for _, getItem := range models.Uni() {
		if getItem.Abbrv == {
			json.NewEncoder(w).Encode(getItem)
		}
	}
}
