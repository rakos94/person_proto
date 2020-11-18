package dao

import "person_proto/models"

// EducationDaoImpl ...
type EducationDaoImpl struct{}

// CreateEducation ...
func (EducationDaoImpl) CreateEducation(data *models.Education) (*models.Education, error) {
	result := g.Create(&data)
	if result.Error == nil {
		return data, nil
	}
	return nil, result.Error
}

// GetEducationAll ...
func (EducationDaoImpl) GetEducationAll() ([]models.Education, error) {
	data := []models.Education{}
	result := g.Find(&data)
	if result.Error != nil {
		return data, result.Error
	}

	return data, nil
}

// GetEducationByID ...
func (EducationDaoImpl) GetEducationByID(id string) (models.Education, error) {
	data := models.Education{}
	result := g.Preload("Person").Where("id", id).First(&data)
	if result.Error != nil {
		return data, result.Error
	}

	return data, nil
}

// UpdateEducation ...
func (EducationDaoImpl) UpdateEducation(id string, data *models.Education) (*models.Education, error) {
	m := &models.Education{}
	first := g.Where("id", id).First(m)
	if first.Error != nil {
		return m, first.Error
	}
	update := g.Model(m).Updates(data)
	if update.Error != nil {
		return m, update.Error
	}

	return m, nil
}
