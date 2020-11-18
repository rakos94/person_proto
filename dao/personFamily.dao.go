package dao

import "person_proto/models"

type PersonFamilyFDao interface {
	CreatePersonFamily(data *models.PersonFamily) error
	GetAllPesonFamily() ([]models.PersonFamily, error)
	GetByIdPersonFamily(id string) (models.PersonFamily, error)
	UpdatePersonFamily(id string, data *models.PersonFamily) error
	DeletePersonFamily(id string) error
}
