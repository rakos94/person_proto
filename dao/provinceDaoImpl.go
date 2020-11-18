package dao

import "person_proto/models"

type ProvinceDaoImpl struct{}

func (ProvinceDaoImpl) CreateProvince(data *models.Provinces) error {
	result := g.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (ProvinceDaoImpl) GetProvinceAll() ([]models.Provinces, error) {
	var (
		data []models.Provinces
	)
	result := g.Preload("Country").Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}
func (ProvinceDaoImpl) GetProvinceById(id string) (models.Provinces, error) {
	var (
		data models.Provinces
	)
	result := g.Preload("Country").Where("id", id).First(&data)
	return data, result.Error
}
func (ProvinceDaoImpl) UpdateProvince(id string, data *models.Provinces) error {
	var m = new(models.Provinces)
	first := g.Where("id", id).First(&m)
	if first.Error != nil {
		return first.Error
	}
	update := g.Model(m).Updates(data)
	if update.Error != nil {
		return update.Error
	}
	return nil
}
func (ProvinceDaoImpl) DeleteProvince(id string) error {
	var m = new(models.Provinces)
	first := g.Where("id", id).Delete(&m)
	if first.Error != nil {
		return first.Error
	}
	return nil
}
