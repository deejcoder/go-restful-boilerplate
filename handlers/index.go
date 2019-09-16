package handlers

import (
	"net/http"

	"github.com/deejcoder/go-restful-boilerplate/handlers/helpers"
)

// Index index
func Index(w http.ResponseWriter, r *http.Request) {
	response := helpers.NewResponse()
	validated := helpers.ValidateClient(r)

	if !validated {
		response.Error(w, "Client not yet authorized", helpers.ErrorNotAuthorized)
	}
}
