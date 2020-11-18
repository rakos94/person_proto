package services

import "person_proto/models"

// AddressService ...
type AddressService interface {
	CreateAddress(data *models.Address) (*models.Address, error)
	GetAddressAll() ([]models.Address, error)
	GetAddressByID(id string) (models.Address, error)
	GetAddressByPersonID(id string) ([]models.Address, error)
	UpdateAddress(id string, data *models.Address) (*models.Address, error)
}
