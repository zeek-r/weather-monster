package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	logger "github.com/sirupsen/logrus"
	handler "github.com/zeek-r/weather-monster/api/http/handlers"
	router "github.com/zeek-r/weather-monster/api/http/routes"
	application "github.com/zeek-r/weather-monster/app"
)

type API struct {
	app *application.App
}

func NewAPI(app *application.App) *API {
	return &API{app}
}

func (api *API) Serve() error {

	// Load Config
	config, err := LoadConfig()
	if err != nil {
		logger.Error("Error loading config", err)
		return err
	}

	// Handler CORS
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-ENVAPI-USERNAME"}),
	)

	// Load Handlers
	Handlers := handler.Init(api.app)

	routes := router.LoadRoutes("/", Handlers)
	logger.Info("Loaded api Routes")

	server := &http.Server{
		Addr:        fmt.Sprintf(":%d", config.Port),
		Handler:     cors(routes),
		ReadTimeout: 2 * time.Minute,
	}
	logger.Infof("serving api at %d", config.Port)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Error(err)
		return err
	}
	return nil
}
