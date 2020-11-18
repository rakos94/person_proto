package models

type Country struct {
	BaseModels
	Name        string `gorm:"column:name" json:"name"`
	Code        string `gorm:"column:code;unique" json:"code"`
	Nasionality string `gorm:"column:nasionality" json:"nasionality"`
}

func (Country) TableName() string {
	return "cor_country"
}
