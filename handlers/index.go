package handlers

import (
	"net/http"

	"github.com/deejcoder/go-restful-boilerplate/helpers"
)

// Index index
func Index(w http.ResponseWriter, r *http.Request) {
	response := helpers.NewResponse()
	response.Success(w, "Authorization validated", nil)
}
