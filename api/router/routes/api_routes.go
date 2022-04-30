package routes

import (
	"net/http"

	"github.com/milinches/nigeria-universities-api/api/controllers"
)

type (
	Route struct {
		URL     string
		Method  string
		Handler func(http.ResponseWriter, *http.Request)
	}
)

var apiRoutes = []Route{
	{
		URL:     "/api/v1",
		Method:  http.MethodGet,
		Handler: controllers.GetUniversities,
	},
	{
		URL:     "/api/v1/{abbreviation}",
		Method:  http.MethodGet,
		Handler: controllers.GetSpecificUniversity,
	},
}
