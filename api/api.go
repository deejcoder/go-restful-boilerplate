package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/olivere/elastic"

	"github.com/deejcoder/go-restful-boilerplate/storage"
	"github.com/deejcoder/go-restful-boilerplate/util/config"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// AppContext stores shared application data for use within Requests
type AppContext struct {
	Db     *elastic.Client
	Config *config.Config
}

type appContextKey struct{}

func configure(ac *AppContext) *http.Server {

	config := config.GetConfig()
	cors := handlers.CORS(
		handlers.AllowedOrigins(config.API.AllowedOrigins),
		handlers.AllowedMethods(config.API.AllowedMethods),
		handlers.AllowedHeaders(config.API.AllowedHeaders),
	)

	router := mux.NewRouter()
	BuildRouter(router)

	// configure server
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", ac.Config.API.Port),
		Handler: ControllerWrapper(ac, cors(router)),
	}

	return s
}

// Start starts the webserver, terminates on request
func Start(ctx context.Context) {

	db := storage.Connect()
	appContext := AppContext{
		Db:     db,
		Config: config.GetConfig(),
	}

	server := configure(&appContext)

	// listen for interupt signal
	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			log.Error(err)
		}
		close(done)
	}()

	// start webserver
	log.Infof("Starting REST api on http://localhost:%d", appContext.Config.API.Port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Error(err)
	}

	<-done
}

// GetAppContext returns the AppContext from a given request
func GetAppContext(r *http.Request) *AppContext {
	ac, _ := r.Context().Value(appContextKey{}).(*AppContext)
	return ac
}
