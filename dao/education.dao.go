package dao

import "person_proto/models"

// EducationDao ...
type EducationDao interface {
	CreateEducation(data *models.Education) (*models.Education, error)
	GetEducationAll() ([]models.Education, error)
	GetEducationByID(id string) (models.Education, error)
	UpdateEducation(id string, data *models.Education) (*models.Education, error)
}
