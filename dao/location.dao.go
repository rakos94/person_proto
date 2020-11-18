package dao

import "person_proto/models"

type locationDao interface {
	CreateLocation(data *models.Location) (*models.Location, error)
}
