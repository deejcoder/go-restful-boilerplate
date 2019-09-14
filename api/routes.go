/*
routes seperates routes from the main application logic, define routes or paths in
defineRoutes().
*/

package api

import (
	"net/http"

	controllers "github.com/deejcoder/go-restful-boilerplate/api/controllers"
	"github.com/gorilla/mux"
)

// Route defines a Route
type Route struct {
	Path       string
	Controller func(http.ResponseWriter, *http.Request)
	Methods    string
}

// BuildRouter creates a new Router and adds all defined routes to it
func BuildRouter() *mux.Router {
	router := mux.NewRouter()

	routes := defineRoutes()
	for _, route := range routes {
		router.HandleFunc(route.Path, route.Controller).Methods(route.Methods)
	}
	return router
}

func defineRoutes() []Route {
	return []Route{
		{Path: "/api", Controller: controllers.Index, Methods: "GET"},
	}
}
