package services

import (
	"errors"
	"person_proto/dao"
	"person_proto/models"
)

var educationDao dao.EducationDao = dao.EducationDaoImpl{}

// EducationServiceImpl ...
type EducationServiceImpl struct {
}

// CreateEducation ...
func (EducationServiceImpl) CreateEducation(data *models.Education) (*models.Education, error) {
	_, err := personService.GetPersonByID(data.PersonID)
	if err != nil {
		return nil, errors.New("Person id not exist")
	}
	return educationDao.CreateEducation(data)
}

// GetEducationAll ...
func (EducationServiceImpl) GetEducationAll() ([]models.Education, error) {
	return educationDao.GetEducationAll()
}

// GetEducationByID ...
func (EducationServiceImpl) GetEducationByID(id string) (models.Education, error) {
	return educationDao.GetEducationByID(id)
}

// GetEducationByPersonID ...
func (EducationServiceImpl) GetEducationByPersonID(id string) ([]models.Education, error) {
	_, err := personService.GetPersonByID(id)
	if err != nil {
		return nil, errors.New("Person id not exist")
	}
	return educationDao.GetEducationByPersonID(id)
}

// UpdateEducation ...
func (EducationServiceImpl) UpdateEducation(id string, data *models.Education) (*models.Education, error) {
	return educationDao.UpdateEducation(id, data)
}
