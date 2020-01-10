package temperatureservice

import (
	"encoding/json"
	"time"

	temperaturerepo "github.com/zeek-r/weather-monster/db/temperatureRepo"
	webhookrepo "github.com/zeek-r/weather-monster/db/webhookRepo"
	"github.com/zeek-r/weather-monster/models"
	logger "github.com/zeek-r/weather-monster/pkg/logger"
	"github.com/zeek-r/weather-monster/pkg/request"
)

type TemperatureService struct {
	temperatureRepo temperaturerepo.Repository
	webhookRepo     webhookrepo.Repository
}

func NewTemperatureService(temperature temperaturerepo.Repository, webhook webhookrepo.Repository) Service {
	return &TemperatureService{
		temperatureRepo: temperature,
		webhookRepo:     webhook,
	}
}

func (service *TemperatureService) Create(toInsert *models.Temperature) error {
	toInsert.Timestamp = time.Now().Unix()
	err := service.temperatureRepo.Create(toInsert)
	go service.TriggerCallback(toInsert)
	return err
}

func (service *TemperatureService) TriggerCallback(temperature *models.Temperature) {
	cityID := temperature.CityID
	webhooks, err := service.webhookRepo.GetWebhooksByID(cityID)

	if err != nil {
		logger.Info("Error getting webhooks related to city", map[string]interface{}{
			"cityID": cityID,
		})
		return
	}
	body, _ := json.Marshal(map[string]interface{}{
		"city_id": temperature.CityID,
		"max":     temperature.Max,
		"min":     temperature.Min,
	})

	for _, webhook := range webhooks {
		logger.Info("Webhook callback trigerred", map[string]interface{}{
			"url":     webhook.CallbackUrl,
			"city_id": temperature.CityID,
		})
		if err := request.Request(body, "POST", webhook.CallbackUrl); err != nil {
			logger.Fatal(err, "Cannot post to webhook")
		}
	}
	return
}
