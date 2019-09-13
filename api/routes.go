package api

import (
	"net/http"

	controllers "github.com/Dilicor/lts/api/controllers"
	"github.com/gorilla/mux"
)

// Route defines a Route
type Route struct {
	Path       string
	Controller func(http.ResponseWriter, *http.Request)
	Methods    string
}

// BuildRouter creates a new Router and adds all defined routes to it
func BuildRouter(router *mux.Router) {
	routes := defineRoutes()
	for _, route := range routes {
		router.HandleFunc(route.Path, route.Controller).Methods(route.Methods)
	}
}

func defineRoutes() []Route {
	return []Route{
		{Path: "/api", Controller: controllers.Index, Methods: "GET"},
	}
}
