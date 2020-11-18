package requests

import (
	"log"
	"person_proto/helper"
	"time"

	"github.com/labstack/echo"
)

// GetValue get value of form
func GetValue(c echo.Context, name string) string {
	return c.FormValue(name)
}

// ToDate convert string to date
func ToDate(snapshot string) helper.Date {
	const layout = "2006-01-02"
	t, err := time.Parse(layout, snapshot)
	if err != nil {
		log.Fatal("Error parse Todate : ", err)
	}
	return helper.Date(t)
}
