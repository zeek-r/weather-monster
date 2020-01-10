package handler

import (
	cityhandler "github.com/zeek-r/weather-monster/api/http/handlers/city"
	forecasthandler "github.com/zeek-r/weather-monster/api/http/handlers/forecast"
	temperaturehandler "github.com/zeek-r/weather-monster/api/http/handlers/temperature"
	webhookhandler "github.com/zeek-r/weather-monster/api/http/handlers/webhook"

	"github.com/zeek-r/weather-monster/app"
)

type Handler struct {
	CityHandler        *cityhandler.CityHandler
	ForecastHandler    *forecasthandler.ForecastHandler
	TemperatureHandler *temperaturehandler.TemperatureHandler
	WebhookHandler     *webhookhandler.WebhookHandler
}

func Init(app *app.App) *Handler {
	cityHandler := cityhandler.NewCityHandler(app.Services.CityService)
	forecastHandler := forecasthandler.NewForecastHandler(app.Services.ForecastService)
	temperatureHandler := temperaturehandler.NewTemperatureHandler(app.Services.TemperatureService)
	webhookHandler := webhookhandler.NewWebhookHandler(app.Services.WebhookService)

	return &Handler{
		CityHandler:        cityHandler,
		ForecastHandler:    forecastHandler,
		TemperatureHandler: temperatureHandler,
		WebhookHandler:     webhookHandler,
	}
}
