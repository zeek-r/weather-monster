package routes

import (
	"github.com/gorilla/mux"
	handler "github.com/zeek-r/weather-monster/api/http/handlers"
	middleware "github.com/zeek-r/weather-monster/api/http/middlewares"
)

func LoadRoutes(pathPrefix string, handlers *handler.Handler) *mux.Router {
	router := mux.NewRouter()

	router = router.PathPrefix(pathPrefix).Subrouter()
	router = router.StrictSlash(true)
	router.Use(middleware.RequestPreProcessor)

	/** Start Route Loading */
	LoadCityRoutes(router, handlers.CityHandler)
	LoadTemperatureRoutes(router, handlers.TemperatureHandler)
	LoadForecastRoutes(router, handlers.ForecastHandler)
	LoadWebhookRoutes(router, handlers.WebhookHandler)
	/** Route Loading End */
	return router
}
