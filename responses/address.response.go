package responses

import (
	"person_proto/helper"
	"person_proto/models"
)

// AddressResponse ...
type AddressResponse struct {
	models.BaseModels
	PersonID   string          `json:"person_id"`
	CityID     string          `json:"city_id"`
	Address    string          `json:"address"`
	PostalCode *string         `json:"postal_code"`
	Phone      *string         `json:"phon"`
	Fax        *string         `json:"fax"`
	IsDefault  bool            `json:"is_default"`
	Person     *PersonResponse `json:"-"`
	models.BaseCUModels
}

// NewAddressResponse ...
func NewAddressResponse(m *models.Address) *AddressResponse {
	return &AddressResponse{
		BaseModels:   m.BaseModels,
		PersonID:     m.PersonID,
		CityID:       m.CityID,
		Address:      m.Address,
		PostalCode:   m.PostalCode,
		Phone:        m.Phone,
		Fax:          m.Fax,
		IsDefault:    m.IsDefault,
		Person:       NewPersonResponse(&m.Person),
		BaseCUModels: m.BaseCUModels,
	}
}

// NewListAddressResponse ...
func NewListAddressResponse(list []models.Address) []AddressResponse {
	var responses []AddressResponse
	for _, d := range list {
		responses = append(responses, *NewAddressResponse(&d))
	}
	return responses
}

// LoadRelation ...
func (r AddressResponse) LoadRelation() map[string]interface{} {
	m := helper.ConvertMap(r)
	m["person"] = r.Person

	return m
}
