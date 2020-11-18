package services

import (
	"fmt"
	"person_proto/dao"
	"person_proto/models"
)

var personFamilyDao dao.PersonFamilyFDao = dao.PersonFamilyDaoImpl{}
var personService PersonService = PersonServiceImpl{}

type PersonFamilyServiceImpl struct{}

func (PersonFamilyServiceImpl) CreatePersonFamily(data *models.PersonFamily) error {
	fmt.Println("service", data)
	_, err := personService.GetPersonByID(data.PersonId)
	if err != nil {
		return err
	}
	return personFamilyDao.CreatePersonFamily(data)
}
func (PersonFamilyServiceImpl) GetAllPesonFamily() ([]models.PersonFamily, error) {
	return personFamilyDao.GetAllPesonFamily()
}

func (PersonFamilyServiceImpl) GetByIdPersonFamily(id string) (models.PersonFamily, error) {
	return personFamilyDao.GetByIdPersonFamily(id)
}

func (PersonFamilyServiceImpl) UpdatePersonFamily(id string, data *models.PersonFamily) error {
	return personFamilyDao.UpdatePersonFamily(id, data)
}
func (PersonFamilyServiceImpl) DeletePersonFamily(id string) error {
	return personFamilyDao.DeletePersonFamily(id)
}
