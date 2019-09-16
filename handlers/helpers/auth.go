package helpers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deejcoder/go-restful-boilerplate/util/config"
	"github.com/dgrijalva/jwt-go"
)

// AuthorizeClient generates a JWT token, and attaches cookie to the ResponseWriter
func AuthorizeClient(w http.ResponseWriter) {
	// generate JWT token
	exp := time.Now().UTC().Add(12 * time.Hour)
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": exp.Unix(),
	})
	token, _ := tk.SignedString([]byte(config.GetConfig().Keys.JWTSecret))

	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/api",
		Expires:  exp,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func ValidateClient(r *http.Request) bool {
	cookie, err := r.Cookie("token")
	if err != nil {
		return false
	}

	tokenValue := cookie.Value

	token, _ := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		// check correct signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(config.GetConfig().Keys.JWTSecret), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}
	return false
}
