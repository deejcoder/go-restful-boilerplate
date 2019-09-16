package api

import (
	"context"
	"fmt"
	"net/http"

	h "github.com/deejcoder/go-restful-boilerplate/api/handlers"
	"github.com/deejcoder/go-restful-boilerplate/storage"
	"github.com/deejcoder/go-restful-boilerplate/util/config"
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
)

func configure(ac *h.AppContext) *http.Server {

	config := config.GetConfig()

	cors := handlers.CORS(
		handlers.AllowedOrigins(config.API.AllowedOrigins),
		handlers.AllowedMethods(config.API.AllowedMethods),
		handlers.AllowedHeaders(config.API.AllowedHeaders),
	)

	router := BuildRouter()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", ac.Config.API.Port),
		Handler: h.HandlerWrapper(ac, cors(router)),
	}

	return s
}

// Start starts the webserver, terminates on request
func Start(ctx context.Context) {

	db := storage.Connect()
	appContext := h.AppContext{
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
