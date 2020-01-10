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
	deleted := repo.Conn.Model(&models.City{}).Delete(toDelete)

	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
