package routes

import (
	"github.com/gorilla/mux"
	"github.com/milinches/nigeria-universities-api/controllers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1", controllers.GetUniversities).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/{abbreviation}", controllers.GetSpecificUniversity).Methods("GET", "OPTIONS")
	return router
}
