package webhookservice_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	webhookservice "github.com/zeek-r/weather-monster/app/services/webhook"
	"github.com/zeek-r/weather-monster/db/webhookRepo/mocks"
	"github.com/zeek-r/weather-monster/models"
)

func TestCreate(t *testing.T) {
	webhookRepoMock := new(mocks.Repository)
	webhookMock := &models.Webhook{
		CityID:      1,
		CallbackUrl: "https://me.me/me",
	}

	t.Run("create success", func(t *testing.T) {
		webhookRepoMock.On("Create", mock.AnythingOfType("*models.Webhook")).Return(nil).Once()
		service := webhookservice.NewWebhookService(webhookRepoMock)
		created := service.Create(webhookMock)
		assert.Equal(t, created, nil)
	})

	t.Run("create error", func(t *testing.T) {
		webhookRepoMock.On("Create", mock.AnythingOfType("*models.Webhook")).Return(errors.New("Creation error")).Once()
		service := webhookservice.NewWebhookService(webhookRepoMock)
		created := service.Create(webhookMock)
		assert.Error(t, created)
	})

}

func TestDeleteByID(t *testing.T) {
	webhookRepoMock := new(mocks.Repository)

	t.Run("DeleteById success", func(t *testing.T) {
		webhookRepoMock.On("DeleteByID", mock.AnythingOfType("*models.Webhook")).Return(nil).Once()
		service := webhookservice.NewWebhookService(webhookRepoMock)
		created := service.DeleteByID(uint(1))
		assert.Equal(t, created, nil)
	})

	t.Run("DeleteById error", func(t *testing.T) {
		webhookRepoMock.On("DeleteByID", mock.AnythingOfType("*models.Webhook")).Return(errors.New("Deletion error")).Once()
		service := webhookservice.NewWebhookService(webhookRepoMock)
		created := service.DeleteByID(uint(1))
		assert.Error(t, created)
	})
}
