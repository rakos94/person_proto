package models

type PersonDocument struct {
	BaseModels
	Name string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	FileDocument string `gorm:"column:file_document" json:"file_document"`
	PersonId string `gorm:"column:person_id" json:"person_id"`
	Person Person `gorm:"foreignKey:person_id" json:"person"`
}

func (PersonDocument) TableName() string {
	return "cor_person_document"
}
