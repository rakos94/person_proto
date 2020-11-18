package services

import (
	"person_proto/dao"
	"person_proto/models"
)

var personFamilyAddressDao dao.PersonFamilyAddress = dao.PersonFamilyAddressDaoImpl{}
var personFamilyService PersonFamilyService = PersonFamilyServiceImpl{}
var cityService CityService = CityServiceImpl{}

type PersonFamilyAddressServiceImpl struct{}

func (PersonFamilyAddressServiceImpl) CreatePersonFamilyAddress(data *models.PersonFamilyAddres) error {
	_, err := personFamilyService.GetByIdPersonFamily(data.PersonFamilyId)
	if err != nil {
		return err
	}
	_, err = cityService.GetCityById(data.CityId)
	if err != nil {
		return err
	}
	return personFamilyAddressDao.CreatePersonFamilyAddress(data)
}

func (PersonFamilyAddressServiceImpl) GetAllPersonFamilyAddess() ([]models.PersonFamilyAddres, error) {
	return personFamilyAddressDao.GetAllPersonFamilyAddess()
}
func (PersonFamilyAddressServiceImpl) GetByIdPersonFamilyAddress(id string) (models.PersonFamilyAddres, error) {
	return personFamilyAddressDao.GetByIdPersonFamilyAddress(id)
}
func (PersonFamilyAddressServiceImpl) UpdatePersonFamilyAddress(id string, data *models.PersonFamilyAddres) error {
	_, err := personFamilyService.GetByIdPersonFamily(data.PersonFamilyId)
	if err != nil {
		return err
	}
	_, err = cityService.GetCityById(data.CityId)
	if err != nil {
		return err
	}
	return personFamilyAddressDao.UpdatePersonFamilyAddress(id, data)
}
func (PersonFamilyAddressServiceImpl) DeletePersonFamilyAddress(id string) error {
	return personFamilyAddressDao.DeletePersonFamilyAddress(id)
}
