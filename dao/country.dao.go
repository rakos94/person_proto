package dao

import "person_proto/models"

type CountryDao interface {
	CreateCountry(data *models.Country) error
	GetCountryAll() ([]models.Country, error)
	GetCountryById(id string) (models.Country, error)
	UpdateCountry(id string, data *models.Country) error
	DeleteCountry(id string) error
}
