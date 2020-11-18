package responses

import (
	"person_proto/helper"
	"person_proto/models"
)

// PersonResponse ...
type PersonResponse struct {
	models.BaseModels
	FirstName         string      `json:"first_name"`
	LastName          *string     `json:"last_name"`
	Email             string      `json:"email"`
	BirthPlace        *string     `json:"birth_place"`
	BirthDate         helper.Date `json:"birth_date"`
	Mobile            *string     `json:"mobile"`
	Gender            string      `json:"gender"`
	IsFromRecruitment bool        `json:"is_from_recruitment"`
	models.BaseCUModels
}

// NewPersonResponse ...
func NewPersonResponse(m *models.Person) *PersonResponse {
	return &PersonResponse{
		BaseModels:        m.BaseModels,
		FirstName:         m.FirstName,
		LastName:          m.LastName,
		Email:             m.Email,
		BirthPlace:        m.BirthPlace,
		BirthDate:         m.BirthDate,
		Mobile:            m.Mobile,
		Gender:            m.Gender,
		IsFromRecruitment: m.IsFromRecruitment,
		BaseCUModels:      m.BaseCUModels,
	}
}

// NewListPersonResponse ...
func NewListPersonResponse(list []models.Person) []PersonResponse {
	var responses []PersonResponse
	for _, d := range list {
		responses = append(responses, *NewPersonResponse(&d))
	}
	return responses
}
