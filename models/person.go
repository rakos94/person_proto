package models

import (
	"person_proto/helper"
)

// Person ...
type Person struct {
	BaseModels
	FirstName         string      `gorm:"not null"`
	LastName          *string     ``
	Email             string      `gorm:"unique;not null"`
	Password          string      `gorm:"not null"`
	BirthPlace        *string     ``
	BirthDate         helper.Date `gorm:"not null;type:date"`
	Mobile            *string     ``
	Gender            string      `gorm:"not null"`
	IsFromRecruitment bool        `gorm:"not null;default:false"`
	CreatedBy         string      `gorm:"not null;default:admin"`
	UpdatedBy         string      `gorm:"not null;default:admin"`
	Token             string      `gorm:"-"`
	Addresses         []Address   `gorm:"foreignKey:PersonID"`
	Educations        []Education `gorm:"foreignKey:PersonID"`
	BaseCUModels
}

// TableName ...
func (Person) TableName() string {
	return "cor_persons"
}
