/*
routes seperates routes from the main application logic, define routes or paths in
defineRoutes().
*/

package api

import (
	"net/http"

	handlers "github.com/deejcoder/go-restful-boilerplate/handlers"
	helpers "github.com/deejcoder/go-restful-boilerplate/helpers"
	"github.com/gorilla/mux"
)

// Route defines a Route
type Route struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
	Methods string
}

// BuildRouter creates a new Router and adds all defined routes to it
func BuildRouter() *mux.Router {
	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()

	routes := defineRoutes()
	for _, route := range routes {
		api.HandleFunc(route.Path, route.Handler).Methods(route.Methods)
	}
	return router
}

func defineRoutes() []Route {
	return []Route{
		{Path: "/", Handler: helpers.RequireAuth(handlers.Index), Methods: "GET"},
		{Path: "/token", Handler: handlers.Authenticate, Methods: "POST"},
	}
}
