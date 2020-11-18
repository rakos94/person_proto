package requests

import (
	"person_proto/models"
	"strconv"

	"github.com/labstack/echo"
)

// AddressUpdateRequest ...
type AddressUpdateRequest struct {
	CityID     string `form:"city_id" validate:"required"`
	Address    string `form:"address" validate:"required"`
	PostalCode string `form:"postal_code"`
	Phone      string `form:"phone"`
	Fax        string `form:"fax"`
	IsDefault  bool   `form:"is_default"`
	UpdatedBy  string `form:"updated_by" validate:"required"`
}

// Convert from echo FormValue
func (r AddressUpdateRequest) Convert(c echo.Context) *AddressUpdateRequest {
	IsDefault, _ := strconv.ParseBool(GetValue(c, "is_default"))
	return &AddressUpdateRequest{
		CityID:     GetValue(c, "city_id"),
		Address:    GetValue(c, "address"),
		PostalCode: GetValue(c, "postal_code"),
		Phone:      GetValue(c, "phone"),
		Fax:        GetValue(c, "fax"),
		IsDefault:  IsDefault,
		UpdatedBy:  GetValue(c, "updated_by"),
	}
}

// Model from request
func (r AddressUpdateRequest) Model() *models.Address {
	return &models.Address{
		CityID:     r.CityID,
		Address:    r.Address,
		PostalCode: &r.PostalCode,
		Phone:      &r.Phone,
		Fax:        &r.Fax,
		IsDefault:  r.IsDefault,
		UpdatedBy:  r.UpdatedBy,
	}
}
