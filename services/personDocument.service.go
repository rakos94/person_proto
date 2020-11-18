package services

import "person_proto/requests"

type PersonDocumentService interface {
	CreatePersonDocumnt(data *requests.PersonDocumntRequest) error
}
