package temperatureservice_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	temperatureService "github.com/zeek-r/weather-monster/app/services/temperature"
	"github.com/zeek-r/weather-monster/db/temperatureRepo/mocks"
	"github.com/zeek-r/weather-monster/models"
)

func TestCreate(t *testing.T) {
	temperatureRepoMock := new(mocks.Repository)
	temperatureMock := &models.Temperature{
		CityID: uint(1),
		Max:    20,
		Min:    18,
	}

	t.Run("create success", func(t *testing.T) {
		temperatureRepoMock.On("Create", mock.AnythingOfType("*models.Temperature")).Return(nil).Once()
		service := temperatureService.NewTemperatureService(temperatureRepoMock)
		created := service.Create(temperatureMock)
		assert.Equal(t, created, nil)
	})

	t.Run("create error", func(t *testing.T) {
		temperatureRepoMock.On("Create", mock.AnythingOfType("*models.Temperature")).Return(errors.New("Error inserting")).Once()
		service := temperatureService.NewTemperatureService(temperatureRepoMock)
		created := service.Create(temperatureMock)
		assert.Error(t, created)
	})

}
