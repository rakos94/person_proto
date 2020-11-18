package requests

import (
	"person_proto/models"
	"strconv"

	"github.com/labstack/echo"
)

// AddressRequest ...
type AddressRequest struct {
	PersonID   string `form:"person_id" validate:"required"`
	CityID     string `form:"city_id" validate:"required"`
	Address    string `form:"address" validate:"required"`
	PostalCode string `form:"postal_code"`
	Phone      string `form:"phone"`
	Fax        string `form:"fax"`
	IsDefault  bool   `form:"is_default"`
	CreatedBy  string `form:"created_by" validate:"required"`
	UpdatedBy  string `form:"updated_by" validate:"required"`
}

// Convert from echo FormValue
func (r AddressRequest) Convert(c echo.Context) *AddressRequest {
	IsDefault, _ := strconv.ParseBool(GetValue(c, "is_default"))
	return &AddressRequest{
		PersonID:   GetValue(c, "person_id"),
		CityID:     GetValue(c, "city_id"),
		Address:    GetValue(c, "address"),
		PostalCode: GetValue(c, "postal_code"),
		Phone:      GetValue(c, "phone"),
		Fax:        GetValue(c, "fax"),
		IsDefault:  IsDefault,
		CreatedBy:  GetValue(c, "created_by"),
		UpdatedBy:  GetValue(c, "updated_by"),
	}
}

// Model from request
func (r AddressRequest) Model() *models.Address {
	return &models.Address{
		PersonID:   r.PersonID,
		CityID:     r.CityID,
		Address:    r.Address,
		PostalCode: &r.PostalCode,
		Phone:      &r.Phone,
		Fax:        &r.Fax,
		IsDefault:  r.IsDefault,
		CreatedBy:  r.CreatedBy,
		UpdatedBy:  r.UpdatedBy,
	}
}
