package cityrepo

import (
	"github.com/jinzhu/gorm"
	"github.com/zeek-r/weather-monster/models"
)

type cityRepository struct {
	Conn *gorm.DB
}

func Init(Conn *gorm.DB) Repository {

	return &cityRepository{Conn}
}

func (repo *cityRepository) Create(toInsert *models.City) error {
	created := repo.Conn.Create(toInsert)
	if created.Error != nil {
		return created.Error
	}
	return nil
}

func (repo *cityRepository) Update(toUpdate *models.City) error {
	updated := repo.Conn.Model(&models.City{}).Updates(toUpdate)

	if updated.Error != nil {
		return updated.Error
	}
	return nil
}

func (repo *cityRepository) DeleteByID(toDelete *models.City) error {
	deleted := repo.Conn.Model(&models.City{}).Delete(toDelete)

	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}

func (repo *cityRepository) GetByID(id int64) (models.City, error) {
	var city models.City
	repo.Conn.First(&city, id)
	return city, nil
}
