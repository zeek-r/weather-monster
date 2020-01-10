package cityservice_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	cityservice "github.com/zeek-r/weather-monster/app/services/city"
	"github.com/zeek-r/weather-monster/db/cityRepo/mocks"
	"github.com/zeek-r/weather-monster/models"
)

func TestCreate(t *testing.T) {
	cityRepoMock := new(mocks.Repository)
	cityMock := &models.City{
		Name:      "Kathmandu",
		Latitude:  "80.8",
		Longitude: "81.88",
	}

	t.Run("create success", func(t *testing.T) {
		cityRepoMock.On("Create", mock.AnythingOfType("*models.City")).Return(nil).Once()
		service := cityservice.NewCityService(cityRepoMock)
		created := service.Create(cityMock)
		assert.Equal(t, created, nil)
	})

	t.Run("create error", func(t *testing.T) {
		cityRepoMock.On("Create", mock.AnythingOfType("*models.City")).Return(errors.New("Unique index error")).Once()
		service := cityservice.NewCityService(cityRepoMock)
		created := service.Create(cityMock)
		assert.Error(t, created)
	})

}

func TestUpdate(t *testing.T) {
	cityRepoMock := new(mocks.Repository)
	cityMock := &models.City{
		Name:      "Kathmandu",
		Latitude:  "80.8",
		Longitude: "81.88",
	}

	t.Run("update success", func(t *testing.T) {
		cityRepoMock.On("Update", mock.AnythingOfType("*models.City")).Return(nil).Once()
		service := cityservice.NewCityService(cityRepoMock)
		created := service.Update(cityMock)
		assert.Equal(t, created, nil)
	})

	t.Run("update error", func(t *testing.T) {
		cityRepoMock.On("Update", mock.AnythingOfType("*models.City")).Return(errors.New("Unique index error")).Once()
		service := cityservice.NewCityService(cityRepoMock)
		created := service.Update(cityMock)
		assert.Error(t, created)
	})
}

func TestDeleteByID(t *testing.T) {
	cityRepoMock := new(mocks.Repository)

	t.Run("DeleteById success", func(t *testing.T) {
		cityRepoMock.On("DeleteByID", mock.AnythingOfType("*models.City")).Return(nil).Once()
		service := cityservice.NewCityService(cityRepoMock)
		created := service.DeleteByID(uint(1))
		assert.Equal(t, created, nil)
	})

	t.Run("DeleteById error", func(t *testing.T) {
		cityRepoMock.On("DeleteByID", mock.AnythingOfType("*models.City")).Return(errors.New("Unique index error")).Once()
		service := cityservice.NewCityService(cityRepoMock)
		created := service.DeleteByID(uint(1))
		assert.Error(t, created)
	})
}
