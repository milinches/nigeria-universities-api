package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/milinches/nigeria-universities-api/models"
)

func GetUniversities(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.Write([]byte("<div><h1>It appears that this page is not available.</h1></div>"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var newUniversity []models.NewUniversities

	for _, uni := range models.JsonToStruct() {
		newUniversity = append(newUniversity, models.NewUniversities{
			Name:        uni.Name,
			Abbrv:       uni.Abbrv,
			WebsiteLink: uni.WebsiteLink,
		})
	}

	_ = json.NewEncoder(w).Encode(newUniversity)
}

func GetSpecificUniversity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, getItem := range models.JsonToStruct() {
		if getItem.Abbrv == params["abbreviation"] {
			json.NewEncoder(w).Encode(getItem)
			return
		}
	}
	w.Write([]byte("Sorry, endpoint not available 😴"))
}
