package services

import (
	"errors"
	"person_proto/dao"
	"person_proto/models"
)

var addressDao dao.AddressDao = dao.AddressDaoImpl{}

// AddressServiceImpl ...
type AddressServiceImpl struct{}

// CreateAddress ...
func (AddressServiceImpl) CreateAddress(data *models.Address) (*models.Address, error) {
	_, err := personService.GetPersonByID(data.PersonID)
	if err != nil {
		return nil, errors.New("Person id not exist")
	}
	_, err = cityService.GetCityById(data.CityID)
	if err != nil {
		return nil, errors.New("City id not exist")
	}
	return addressDao.CreateAddress(data)
}

// GetAddressAll ...
func (AddressServiceImpl) GetAddressAll() ([]models.Address, error) {
	return addressDao.GetAddressAll()
}

// GetAddressByID ...
func (AddressServiceImpl) GetAddressByID(id string) (models.Address, error) {
	return addressDao.GetAddressByID(id)
}

// GetAddressByPersonID ...
func (AddressServiceImpl) GetAddressByPersonID(id string) ([]models.Address, error) {
	_, err := personService.GetPersonByID(id)
	if err != nil {
		return nil, errors.New("Person id not exist")
	}
	return addressDao.GetAddressByPersonID(id)
}

// UpdateAddress ...
func (AddressServiceImpl) UpdateAddress(id string, data *models.Address) (*models.Address, error) {
	_, err := cityService.GetCityById(data.CityID)
	if err != nil {
		return nil, errors.New("City id not exist")
	}
	return addressDao.UpdateAddress(id, data)
}
