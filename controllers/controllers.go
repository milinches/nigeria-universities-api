package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUniversities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var newUniversity []NewUniversities

	for _, uni := range ObjMethod() {
		newUniversity = append(newUniversity, NewUniversities{
			Name:         uni.Name,
			Abbreviation: uni.Abbreviation,
			WebsiteLink:  uni.WebsiteLink,
		})
	}

	_ = json.NewEncoder(w).Encode(newUniversity)
}

func GetSpecificUniversity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, getItem := range ObjMethod() {
		if getItem.Abbreviation == params["abbreviation"] {
			json.NewEncoder(w).Encode(getItem)
			return
		}
	}
	w.Write([]byte("Sorry, endpoint not available."))
}
