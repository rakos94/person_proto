package models

type PersonFamily struct {
	BaseModels
	Name string  `gorm:"column:name" json:"name"`
	BirthPlace string  `gorm:"column:birth_place" json:"birth_place"`
	BirthDate string  `gorm:"column:birth_date" json:"birth_date"`
	Phone string  `gorm:"column:phone" json:"phone"`
	PersonId string  `gorm:"column:person_id" json:"person_id"`
	Person Person `gorm:"foreignKey:person_id" json:"person"`
}
func (PersonFamily) TableName() string {
	return "cor_person_families"
}
