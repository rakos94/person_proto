package dao

import "person_proto/models"

type PersonDocumentDao interface {
	CreatePersonDocumnt(data *models.PersonDocument) error
}
