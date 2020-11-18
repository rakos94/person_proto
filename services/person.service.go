package services

import "person_proto/models"

// PersonService ...
type PersonService interface {
	Login(email string, pwd string) (models.Person, error)
	CreatePerson(person *models.Person) (*models.Person, error)
	GetPersonAll() ([]models.Person, error)
	GetPersonByID(id string) (models.Person, error)
	UpdatePerson(id string, person *models.Person) (*models.Person, error)
	UpdatePassword(id string, password string) error
}
