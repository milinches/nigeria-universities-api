package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	models "github.com/milinches/nigeria-universities-api/api/models"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	hello := map[string]string {
		"status": "200",
		"title": "nigeria universities api",
		"msg": "help improve this project by open a PR, thanks. 🥰",
	}
	marshallObj, _ := json.MarshalIndent(hello, "", "	")

	w.Write(marshallObj)
} 

func GetUniversities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var newUniversity []models.NewUniversities

	for _, uni := range models.ObjMethod() {
		newUniversity = append(newUniversity, models.NewUniversities{
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

	for _, getItem := range models.ObjMethod() {
		if strings.ToLower(getItem.Abbreviation) == strings.ToLower(params["abbreviation"]) {
			json.NewEncoder(w).Encode(getItem)
			return
		}
	}
	w.Write([]byte("Sorry, endpoint not available."))
}
