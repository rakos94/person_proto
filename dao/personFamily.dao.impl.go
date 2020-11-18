package dao

import (
	"person_proto/models"
)

type PersonFamilyDaoImpl struct{}

func (PersonFamilyDaoImpl) CreatePersonFamily(data *models.PersonFamily) error {
	result := g.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (PersonFamilyDaoImpl) GetAllPesonFamily() ([]models.PersonFamily, error) {
	var (
		data []models.PersonFamily
	)
	result := g.Preload("Person").Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}
func (PersonFamilyDaoImpl) GetByIdPersonFamily(id string) (models.PersonFamily, error) {
	var c models.PersonFamily
	result := g.Preload("Person").Where("id", id).First(&c)

	return c, result.Error
}

func (PersonFamilyDaoImpl) UpdatePersonFamily(id string, data *models.PersonFamily) error {
	var m = new(models.PersonFamily)
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
func (PersonFamilyDaoImpl) DeletePersonFamily(id string) error {
	var m = new(models.PersonFamily)
	first := g.Where("id", id).Delete(&m)
	if first.Error != nil {
		return first.Error
	}
	return nil
}
