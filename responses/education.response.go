package responses

import (
	"person_proto/helper"
	"person_proto/models"
)

// EducationResponse ...
type EducationResponse struct {
	models.BaseModels
	PersonID    string          `json:"person_id"`
	Institution string          `json:"institution"`
	Subject     string          `json:"subject"`
	Grade       *float64        `json:"grade"`
	YearBegin   *int            `json:"year_begin"`
	YearEnd     *int            `json:"year_end"`
	Person      *PersonResponse `json:"-"`
	models.BaseCUModels
}

// NewEducationResponse ...
func NewEducationResponse(m *models.Education) *EducationResponse {
	return &EducationResponse{
		BaseModels:   m.BaseModels,
		PersonID:     m.PersonID,
		Institution:  m.Institution,
		Subject:      m.Subject,
		Grade:        m.Grade,
		YearBegin:    m.YearBegin,
		YearEnd:      m.YearEnd,
		Person:       NewPersonResponse(&m.Person),
		BaseCUModels: m.BaseCUModels,
	}
}

// NewListEducationResponse ...
func NewListEducationResponse(list []models.Education) []EducationResponse {
	var responses []EducationResponse
	for _, d := range list {
		responses = append(responses, *NewEducationResponse(&d))
	}
	return responses
}

// LoadRelation ...
func (r EducationResponse) LoadRelation() map[string]interface{} {
	m := helper.ConvertMap(r)
	m["person"] = r.Person

	return m
}
