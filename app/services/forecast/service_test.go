package forecastservice_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	forecastservice "github.com/zeek-r/weather-monster/app/services/forecast"
	"github.com/zeek-r/weather-monster/db/temperatureRepo/mocks"
	"github.com/zeek-r/weather-monster/models"
)

func TestGetByID(t *testing.T) {
	forecastRepoMock := new(mocks.Repository)
	forecastMock := &models.Forecast{
		Max:    20.0,
		Min:    19.9,
		Sample: int64(22),
	}
	t.Run("GetForecastByID success", func(t *testing.T) {
		forecastRepoMock.On("GetForecastByID", mock.AnythingOfType("uint")).Return(forecastMock, nil).Once()
		service := forecastservice.NewForecastService(forecastRepoMock)
		forecast, err := service.GetByID(uint(1))
		assert.Empty(t, err)
		assert.Equal(t, forecast.Max, 20.0)
		assert.Equal(t, forecast.Min, 19.9)
		assert.Equal(t, forecast.Sample, int64(22))
	})

	t.Run("GetForecastByID error", func(t *testing.T) {
		forecastRepoMock.On("GetForecastByID", mock.AnythingOfType("uint")).Return(nil, errors.New("Unique index error")).Once()
		service := forecastservice.NewForecastService(forecastRepoMock)
		forecast, err := service.GetByID(uint(1))
		assert.Error(t, err)
		assert.Empty(t, forecast)
	})

}
