package models

type City struct {
	BaseModels
	Name       string    `gorm:"column:name" json:"name"`
	Code       string    `gorm:"column:code;unique" json:"code"`
	ProvinceId string    `gorm:"column:province_id" json:"province_id"`
	Provinces  Provinces `gorm:"foreignKey:province_id" json:"provinces"`
}

func (City) TableName() string {
	return "cor_city"
}
