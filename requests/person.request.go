package requests

import (
	"person_proto/helper"
	"person_proto/models"
	"strconv"

	"github.com/labstack/echo"
)

// PersonRequest ...
type PersonRequest struct {
	FirstName         string      `form:"first_name" validate:"required"`
	LastName          string      `form:"last_name"`
	Email             string      `form:"email" validate:"required,email"`
	Password          string      `form:"password" validate:"required"`
	BirthPlace        string      `form:"birth_place"`
	BirthDate         helper.Date `form:"birth_date" validate:"required"`
	Mobile            string      `form:"mobile"`
	Gender            string      `form:"gender" validate:"required"`
	IsFromRecruitment bool        `form:"is_from_recruitment"`
	CreatedBy         string      `form:"created_by" validate:"required"`
	UpdatedBy         string      `form:"updated_by" validate:"required"`
}

// Convert from echo FormValue
func (p PersonRequest) Convert(c echo.Context) *PersonRequest {
	BirthDate := GetValue(c, "birth_date")
	IsFromRecruitment, _ := strconv.ParseBool(GetValue(c, "is_from_recruitment"))
	return &PersonRequest{
		FirstName:         GetValue(c, "first_name"),
		LastName:          GetValue(c, "last_name"),
		Email:             GetValue(c, "email"),
		Password:          GetValue(c, "password"),
		BirthPlace:        GetValue(c, "birth_place"),
		BirthDate:         ToDate(BirthDate),
		Mobile:            GetValue(c, "mobile"),
		Gender:            GetValue(c, "gender"),
		IsFromRecruitment: IsFromRecruitment,
		CreatedBy:         GetValue(c, "created_by"),
		UpdatedBy:         GetValue(c, "updated_by"),
	}
}

// Model from request
func (p PersonRequest) Model() *models.Person {
	return &models.Person{
		FirstName:         p.FirstName,
		LastName:          &p.LastName,
		Email:             p.Email,
		Password:          p.Password,
		BirthPlace:        &p.BirthPlace,
		BirthDate:         p.BirthDate,
		Mobile:            &p.Mobile,
		Gender:            p.Gender,
		IsFromRecruitment: p.IsFromRecruitment,
		CreatedBy:         p.CreatedBy,
		UpdatedBy:         p.UpdatedBy,
	}
}
