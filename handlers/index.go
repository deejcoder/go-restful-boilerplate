package handlers

import (
	"net/http"

	"github.com/deejcoder/go-restful-boilerplate/handlers/helpers"
)

// Index index
func Index(w http.ResponseWriter, r *http.Request) {
	validated := helpers.ValidateClient(r)
	if !validated {
		http.Error(w, "Not authorized", http.StatusForbidden)
	}
}
