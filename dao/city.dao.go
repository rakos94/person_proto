package dao

import "person_proto/models"

type CityDao interface {
	CreateCity(data *models.City) error
	GetCityAll() ([]models.City, error)
	GetCityById(id string) (models.City, error)
	UpdateCity(id string, data *models.City) error
	DeleteCity(id string) error
}
