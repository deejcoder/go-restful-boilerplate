package handlers

import (
	"net/http"

	"github.com/deejcoder/go-restful-boilerplate/handlers/helpers"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	helpers.AuthorizeClient(w)
}
