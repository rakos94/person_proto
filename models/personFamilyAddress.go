package models

type PersonFamilyAddres struct {
	BaseModels
	Address string `gorm:"column:address" json:"address"`
	PostalCode string `gorm:"column:postal_code" json:"postal_code"`
	Fax string `gorm:"column:fax" json:"fax"`
	MapLocation string `gorm:"column:map_location" json:"map_location"`
	PersonFamilyId string `gorm:"column:person_family_id" json:"person_family_id"`
	PersonFamily PersonFamily `gorm:"foreignKey:person_family_id" json:"person_family"`
	CityId string `gorm:"column:city_id" json:"city_id"`
	City City `gorm:"foreignKey:city_id" json:"city"`
}
func (PersonFamilyAddres) TableName() string {
	return "cor_person_family_address"
}