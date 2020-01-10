package routes

import (
	"github.com/gorilla/mux"
	handler "github.com/zeek-r/weather-monster/api/http/handlers"
	middleware "github.com/zeek-r/weather-monster/api/http/middlewares"
)

func LoadRoutes(pathPrefix string, handler *handler.Handler) *mux.Router {
	router := mux.NewRouter()

	router = router.PathPrefix(pathPrefix).Subrouter()
	router = router.StrictSlash(true)
	router.Use(middleware.RequestPreProcessor)

	/** Start Route Loading */
	LoadCityRoutes(router, handler.CityHandler)
	LoadTemperatureRoutes(router, handler.TemperatureHandler)
	LoadForecastRoutes(router, handler.ForecastHandler)
	LoadWebhookRoutes(router, handler.WebhookHandler)
	/** Route Loading End */
	return router
}
