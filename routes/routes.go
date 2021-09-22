package routes

import (
	"github.com/gorilla/mux"
	"github.com/milinches/nigeria-universities-api/controllers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", controllers.GetUniversities).Methods("GET")
	router.HandleFunc("/api{abbreviation}", controllers.GetSpecificUniversity).Methods("GET")
	return router
}
