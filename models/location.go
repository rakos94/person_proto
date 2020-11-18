package models

// Location ...
type Location struct {
	BaseModels
	Code        string `gorm:"column:code;unique" json:"code"`
	Name        string `gorm:"column:name" json:"name"`
	Address     string `gorm:"column:address" json:"address"`
	PostalCode  string `gorm:"column:postal_code" json:"postal_code"`
	Phone       string `gorm:"column:phone" json:"phone"`
	Description string `gorm:"column:phone" json:"description"`
	CityId      string `gorm:"column:city_id" json:"city_id"`
	City        City   `gorm:"foreignKey:city_id" json:"city"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// TableName ...
func (Location) TableName() string {
	return "tb_locations"
}
