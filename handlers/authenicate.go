package handlers

import (
	"net/http"

	"github.com/deejcoder/go-restful-boilerplate/handlers/helpers"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	helpers.AuthorizeClient(w)

	response := helpers.NewResponse()
	response.Success(w, "Rights authorized", nil)
}
