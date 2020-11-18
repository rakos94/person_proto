package requests

import (
	"person_proto/models"
	"strconv"

	"github.com/labstack/echo"
)

// EducationRequest ...
type EducationRequest struct {
	PersonID    string  `form:"person_id" validate:"required"`
	Institution string  `form:"institution" validate:"required"`
	Subject     string  `form:"subject" validate:"required"`
	Grade       float64 `form:"grade" validate:"required"`
	YearBegin   int     `form:"year_begin"`
	YearEnd     int     `form:"year_end"`
	CreatedBy   string  `form:"created_by" validate:"required"`
	UpdatedBy   string  `form:"updated_by" validate:"required"`
}

// Convert from echo FormValue
func (r EducationRequest) Convert(c echo.Context) *EducationRequest {
	Grade, _ := strconv.ParseFloat(GetValue(c, "grade"), 64)
	YearBegin, _ := strconv.Atoi(GetValue(c, "year_begin"))
	YearEnd, _ := strconv.Atoi(GetValue(c, "year_end"))
	return &EducationRequest{
		PersonID:    GetValue(c, "person_id"),
		Institution: GetValue(c, "institution"),
		Subject:     GetValue(c, "subject"),
		Grade:       Grade,
		YearBegin:   YearBegin,
		YearEnd:     YearEnd,
		CreatedBy:   GetValue(c, "created_by"),
		UpdatedBy:   GetValue(c, "updated_by"),
	}
}

// Model from request
func (r EducationRequest) Model() *models.Education {
	return &models.Education{
		PersonID:    r.PersonID,
		Institution: r.Institution,
		Subject:     r.Subject,
		Grade:       &r.Grade,
		YearBegin:   &r.YearBegin,
		YearEnd:     &r.YearEnd,
		CreatedBy:   r.CreatedBy,
		UpdatedBy:   r.UpdatedBy,
	}
}
