package models

// Address ...
type Address struct {
	BaseModels
	PersonID   string  `gorm:"not null"`
	CityID     string  `gorm:"not null"`
	Address    string  `gorm:"not null"`
	PostalCode *string ``
	Phone      *string ``
	Fax        *string ``
	IsDefault  bool    `gorm:"not null"`
	CreatedBy  string  `gorm:"not null;default:admin"`
	UpdatedBy  string  `gorm:"not null;default:admin"`
	Person     Person  `gorm:"foreignKey:PersonID"`
	BaseCUModels
}

// TableName ...
func (Address) TableName() string {
	return "cor_person_addresses"
}
