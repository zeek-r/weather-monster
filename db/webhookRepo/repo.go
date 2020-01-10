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

	return nil
}

func (repo *webhookRepository) Update(toUpdate *models.Webhook) error {

	return nil
}

func (repo *webhookRepository) DeleteByID(id int64) error {
	return nil
}

func (repo *webhookRepository) GetByID(id int64) (*models.Webhook, error) {
	return nil, nil
}
