package services

import (
	"person_proto/dao"
	"person_proto/models"
)

var provinceDao dao.ProvinceDao = dao.ProvinceDaoImpl{}
var countryService CountryService = CountryServiceImpl{}

type ProvinceServiceImpl struct{}

func (ProvinceServiceImpl) CreateProvince(data *models.Provinces) error {
	_, err := countryService.GetCountryById(data.CountryId)
	if err != nil {
		return err
	}
	return provinceDao.CreateProvince(data)
}
func (ProvinceServiceImpl) GetProvinceAll() ([]models.Provinces, error) {
	return provinceDao.GetProvinceAll()
}
func (ProvinceServiceImpl) GetProvinceById(id string) (models.Provinces, error) {
	return provinceDao.GetProvinceById(id)
}
func (ProvinceServiceImpl) UpdateProvince(id string, data *models.Provinces) error {
	_, err := provinceDao.GetProvinceById(id)
	if err != nil {
		return err
	}
	return provinceDao.UpdateProvince(id, data)
}
func (ProvinceServiceImpl) DeleteProvince(id string) error {
	_, err := provinceDao.GetProvinceById(id)
	if err != nil {
		return err
	}
	return provinceDao.DeleteProvince(id)
}
