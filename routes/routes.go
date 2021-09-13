package routes

import (
	"github.com/milinches/nigeria-universities-api/controllers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.GetUniversities).Methods("GET")
	router.HandleFunc("/{abbreviation}", controllers.GetSpecificUniversity).Methods("GET")
	return router	
}