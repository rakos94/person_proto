package dao

import "person_proto/models"

// PersonDao ...
type PersonDao interface {
	CreatePerson(data *models.Person) (*models.Person, error)
	GetPersonAll() ([]models.Person, error)
	GetPersonByID(id string) (models.Person, error)
	GetPersonByEmail(email string) (models.Person, error)
	UpdatePerson(id string, data *models.Person) (*models.Person, error)
	UpdatePassword(id string, password string) error
}
