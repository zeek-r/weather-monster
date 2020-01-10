package service

import (
	"github.com/jinzhu/gorm"

	cityservice "github.com/zeek-r/weather-monster/app/services/city"
	cityrepo "github.com/zeek-r/weather-monster/db/cityRepo"

	forecastservice "github.com/zeek-r/weather-monster/app/services/forecast"

	temperatureservice "github.com/zeek-r/weather-monster/app/services/temperature"
	temperaturerepo "github.com/zeek-r/weather-monster/db/temperatureRepo"

	webhookservice "github.com/zeek-r/weather-monster/app/services/webhook"
	webhookrepo "github.com/zeek-r/weather-monster/db/webhookRepo"
)

type Service struct {
	CityService        cityservice.Service
	ForecastService    forecastservice.Service
	TemperatureService temperatureservice.Service
	WebhookService     webhookservice.Service
}

func Init(dbConn *gorm.DB) Service {

	/** Repo and services initialization start **/

	cityRepo := cityrepo.Init(dbConn)
	cityService := cityservice.NewCityService(cityRepo)

	webhookRepo := webhookrepo.Init(dbConn)
	webhookService := webhookservice.NewWebhookService(webhookRepo)

	temperatureRepo := temperaturerepo.Init(dbConn)
	temperatureService := temperatureservice.NewTemperatureService(temperatureRepo, webhookRepo)

	forecastService := forecastservice.NewForecastService(temperatureRepo)

	/** Repo and services initialization end **/

	return Service{
		CityService:        cityService,
		ForecastService:    forecastService,
		TemperatureService: temperatureService,
		WebhookService:     webhookService,
	}
}
