package dao

import "person_proto/models"

type CityDaoImpl struct{}

func (CityDaoImpl) CreateCity(data *models.City) error {
	result := g.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (CityDaoImpl) GetCityAll() ([]models.City, error) {
	var c []models.City
	result := g.Preload("Provinces.Country").Find(&c)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}
func (CityDaoImpl) GetCityById(id string) (models.City, error) {
	var c models.City
	result := g.Preload("Provinces.Country").Where("id", id).First(&c)

	return c, result.Error
}
func (CityDaoImpl) UpdateCity(id string, data *models.City) error {
	var m = new(models.City)
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
func (CityDaoImpl) DeleteCity(id string) error {
	var m = new(models.City)
	first := g.Where("id", id).Delete(&m)
	if first.Error != nil {
		return first.Error
	}
	return nil
}
