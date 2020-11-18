package dao

import (
	"person_proto/models"
)

type CountryDaoImpl struct{}

func (CountryDaoImpl) CreateCountry(data *models.Country) error {
	result := g.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (CountryDaoImpl) GetCountryAll() ([]models.Country, error) {
	var c []models.Country
	result := g.Find(&c)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}
func (CountryDaoImpl) GetCountryById(id string) (models.Country, error) {
	var c models.Country
	result := g.Where("id", id).First(&c)

	return c, result.Error

}
func (CountryDaoImpl) UpdateCountry(id string, data *models.Country) error {
	var m = new(models.Country)
	first := g.Where("id", id).First(&m)
	if first.Error != nil {
		return first.Error
	}
	update := g.Model(m).Updates(data)
	if update.Error != nil {
		return update.Error
	}

	return nil
}
func (CountryDaoImpl) DeleteCountry(id string) error {
	var m = new(models.Country)
	first := g.Where("id", id).Delete(&m)
	if first.Error != nil {
		return first.Error
	}
	return nil
}
