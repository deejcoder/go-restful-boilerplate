/*
controller_wrapper encapsulates controllers (http.Handler), to provide
accress to the database client, to enable logging and set generic response parameters such as headers.
*/

package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/deejcoder/go-restful-boilerplate/util/config"
	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
)

// AppContext allows access to shared app data within controllers
type AppContext struct {
	Db     *elastic.Client
	Config *config.Config
}

type appContextKey struct{}

// GetAppContext returns the AppContext from a given request
func GetAppContext(r *http.Request) *AppContext {
	ac, _ := r.Context().Value(appContextKey{}).(*AppContext)
	return ac
}

// ControllerWrapper is a wrapper which wraps all handler functions
func ControllerWrapper(appContext *AppContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// simple logging
		startTime := time.Now()
		defer func() {
			log.WithFields(log.Fields{
				"remote":   r.RemoteAddr,
				"duration": time.Since(startTime),
			}).Infof("%s %s", r.Method, r.URL.RequestURI())

		}()

		w.Header().Set("Content-Type", "application/json")

		// add AppContext to request context
		ctx := context.WithValue(r.Context(), appContextKey{}, appContext)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
