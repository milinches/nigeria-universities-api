package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/milinches/nigeria-universities-api/models"
)

func Uni() []*models.Universities {
	return []*models.Universities {
		&models.Universities{Name: "Abia State University", Abbrv: "ABSU", WebsiteLink: "https://absu.com"},
		&models.Universities{Name: "Abia State University", Abbrv: "ABSU", WebsiteLink: "https://absu.com"},
	}
}

func GetUniversities(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(401)
		w.Write([]byte("Method not Allowed"))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Uni())
}