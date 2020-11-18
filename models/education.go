package models

// Education ...
type Education struct {
	BaseModels
	PersonID    string   `gorm:"not null"`
	Institution string   `gorm:"not null"`
	Subject     string   `gorm:"not null"`
	Grade       *float64 `gorm:"not null"`
	YearBegin   *int     ``
	YearEnd     *int     ``
	CreatedBy   string   `gorm:"not null;default:admin"`
	UpdatedBy   string   `gorm:"not null;default:admin"`
	Person      Person   `gorm:"foreignKey:PersonID"`
	BaseCUModels
}

// TableName ...
func (Education) TableName() string {
	return "cor_person_educations"
}
