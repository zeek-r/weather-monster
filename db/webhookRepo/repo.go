package webhookrepo

import (
	"github.com/jinzhu/gorm"
	"github.com/zeek-r/weather-monster/models"
)

type webhookRepository struct {
	Conn *gorm.DB
}

func Init(Conn *gorm.DB) Repository {

	return &webhookRepository{Conn}
}

func (repo *webhookRepository) Create(toInsert *models.Webhook) error {
	created := repo.Conn.Create(toInsert)
	if created.Error != nil {
		return created.Error
	}
	return nil
}

func (repo *webhookRepository) DeleteByID(toDelete *models.Webhook) error {
	deleted := repo.Conn.Model(&models.Webhook{}).Delete(toDelete)

	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}

func (repo *webhookRepository) GetWebhooksByID(id uint) ([]models.Webhook, error) {
	var webhooks []models.Webhook
	query := "SELECT * FROM webhooks where city_id = ?"
	repo.Conn.Raw(query, id).Scan(&webhooks)
	return webhooks, nil
}
