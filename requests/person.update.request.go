package requests

import (
	"person_proto/helper"
	"person_proto/models"
	"strconv"

	"github.com/labstack/echo"
)

// PersonUpdateRequest ...
type PersonUpdateRequest struct {
	FirstName         string      `form:"first_name"`
	LastName          string      `form:"last_name"`
	Email             string      `form:"email" validate:"email"`
	BirthPlace        string      `form:"birth_place"`
	BirthDate         helper.Date `form:"birth_date"`
	Mobile            string      `form:"mobile"`
	Gender            string      `form:"gender"`
	IsFromRecruitment bool        `form:"is_from_recruitment"`
	UpdatedBy         string      `form:"updated_by"`
}

// Convert from echo FormValue
func (p PersonUpdateRequest) Convert(c echo.Context) *PersonUpdateRequest {
	BirthDate := GetValue(c, "birth_date")
	IsFromRecruitment, _ := strconv.ParseBool(GetValue(c, "is_from_recruitment"))
	return &PersonUpdateRequest{
		FirstName:         GetValue(c, "first_name"),
		LastName:          GetValue(c, "last_name"),
		Email:             GetValue(c, "email"),
		BirthPlace:        GetValue(c, "birth_place"),
		BirthDate:         ToDate(BirthDate),
		Mobile:            GetValue(c, "mobile"),
		Gender:            GetValue(c, "gender"),
		IsFromRecruitment: IsFromRecruitment,
		UpdatedBy:         GetValue(c, "updated_by"),
	}
}

// Model from request
func (p PersonUpdateRequest) Model() *models.Person {
	return &models.Person{
		FirstName:         p.FirstName,
		LastName:          &p.LastName,
		Email:             p.Email,
		BirthPlace:        &p.BirthPlace,
		BirthDate:         p.BirthDate,
		Mobile:            &p.Mobile,
		Gender:            p.Gender,
		IsFromRecruitment: p.IsFromRecruitment,
		UpdatedBy:         p.UpdatedBy,
	}
}
