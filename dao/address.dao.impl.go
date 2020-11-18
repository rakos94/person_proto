package dao

import "person_proto/models"

// AddressDaoImpl ...
type AddressDaoImpl struct{}

// CreateAddress ...
func (AddressDaoImpl) CreateAddress(data *models.Address) (*models.Address, error) {
	result := g.Create(&data)
	if result.Error == nil {
		return data, nil
	}
	return nil, result.Error
}

// GetAddressAll ...
func (AddressDaoImpl) GetAddressAll() ([]models.Address, error) {
	data := []models.Address{}
	result := g.Find(&data)
	if result.Error != nil {
		return data, result.Error
	}

	return data, nil
}

// GetAddressByID ...
func (AddressDaoImpl) GetAddressByID(id string) (models.Address, error) {
	data := models.Address{}
	result := g.Preload("Person").Where("id", id).First(&data)
	if result.Error != nil {
		return data, result.Error
	}

	return data, nil
}

// GetAddressByPersonID ...
func (AddressDaoImpl) GetAddressByPersonID(id string) ([]models.Address, error) {
	data := []models.Address{}
	result := g.Preload("Person").Where("person_id", id).Find(&data)
	if result.Error != nil {
		return data, result.Error
	}

	return data, nil
}

// UpdateAddress ...
func (AddressDaoImpl) UpdateAddress(id string, data *models.Address) (*models.Address, error) {
	m := &models.Address{}
	first := g.Where("id", id).First(m)
	if first.Error != nil {
		return m, first.Error
	}
	update := g.Model(m).Updates(data)
	if update.Error != nil {
		return m, update.Error
	}

	return m, nil
}
