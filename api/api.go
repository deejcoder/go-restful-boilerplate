package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deejcoder/go-restful-boilerplate/helpers"
	"github.com/deejcoder/go-restful-boilerplate/storage"
	"github.com/deejcoder/go-restful-boilerplate/util/config"
	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
)

func configure(ac *helpers.AppContext) *http.Server {

	config := config.GetConfig()

	cors := handlers.CORS(
		handlers.AllowedOrigins(config.API.AllowedOrigins),
		handlers.AllowedMethods(config.API.AllowedMethods),
		handlers.AllowedHeaders(config.API.AllowedHeaders),
	)

	router := BuildRouter()

	// enable csrf tokens
	csrfMiddleware := csrf.Protect(
		[]byte(config.Keys.CSRFKey),
		csrf.Secure(config.API.UsingHttps),
	)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", ac.Config.API.Port),
		Handler: csrfMiddleware(helpers.HandlerWrapper(ac, cors(router))),
	}

	return s
}

// Start starts the webserver, terminates on request
func Start(ctx context.Context) {

	db := storage.Connect()
	appContext := helpers.AppContext{
		Db:     db,
		Config: config.GetConfig(),
	}

	server := configure(&appContext)

	// listen for interupt signal to close server
	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			log.Error(err)
		}
		close(done)
	}()

	log.Infof("Starting REST api on http://localhost:%d", appContext.Config.API.Port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Error(err)
	}

	<-done
}
