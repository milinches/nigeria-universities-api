package routes

import (
	"github.com/gorilla/mux"
)

func Load() []Route {
	routes := apiRoutes
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.URL, route.Handler).Methods(route.Method)
	}
	return r
}
