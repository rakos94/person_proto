package dao

import (
	"person_proto/models"
)

type PersonFamilyAddressDaoImpl struct{}

func (PersonFamilyAddressDaoImpl) CreatePersonFamilyAddress(data *models.PersonFamilyAddres) error {
	result := g.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (PersonFamilyAddressDaoImpl) GetAllPersonFamilyAddess() ([]models.PersonFamilyAddres, error) {
	var (
		data []models.PersonFamilyAddres
	)
	result := g.Preload("City.Provinces.Country").Preload("PersonFamily.Person").Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}
func (PersonFamilyAddressDaoImpl) GetByIdPersonFamilyAddress(id string) (models.PersonFamilyAddres, error) {
	var (
		data models.PersonFamilyAddres
	)
	result := g.Preload("City.Provinces.Country").Preload("PersonFamily.Person").Where("id", id).First(&data)

	return data, result.Error
}

func (PersonFamilyAddressDaoImpl) UpdatePersonFamilyAddress(id string, data *models.PersonFamilyAddres) error {
	var m = new(models.PersonFamilyAddres)
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
func (PersonFamilyAddressDaoImpl) DeletePersonFamilyAddress(id string) error {
	var m = new(models.PersonFamilyAddres)
	first := g.Where("id", id).Delete(&m)
	if first.Error != nil {
		return first.Error
	}
	return nil
}
