package requests

import (
	"person_proto/models"
	"strconv"

	"github.com/labstack/echo"
)

// EducationUpdateRequest ...
type EducationUpdateRequest struct {
	Institution string  `form:"institution" validate:"required"`
	Subject     string  `form:"subject" validate:"required"`
	Grade       float64 `form:"grade" validate:"required"`
	YearBegin   int     `form:"year_begin"`
	YearEnd     int     `form:"year_end"`
	UpdatedBy   string  `form:"updated_by" validate:"required"`
}

// Convert from echo FormValue
func (r EducationUpdateRequest) Convert(c echo.Context) *EducationUpdateRequest {
	Grade, _ := strconv.ParseFloat(GetValue(c, "grade"), 64)
	YearBegin, _ := strconv.Atoi(GetValue(c, "year_begin"))
	YearEnd, _ := strconv.Atoi(GetValue(c, "year_end"))
	return &EducationUpdateRequest{
		Institution: GetValue(c, "institution"),
		Subject:     GetValue(c, "subject"),
		Grade:       Grade,
		YearBegin:   YearBegin,
		YearEnd:     YearEnd,
		UpdatedBy:   GetValue(c, "updated_by"),
	}
}

// Model from request
func (r EducationUpdateRequest) Model() *models.Education {
	return &models.Education{
		Institution: r.Institution,
		Subject:     r.Subject,
		Grade:       &r.Grade,
		YearBegin:   &r.YearBegin,
		YearEnd:     &r.YearEnd,
		UpdatedBy:   r.UpdatedBy,
	}
}
