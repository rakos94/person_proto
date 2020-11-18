package services

import (
	"person_proto/dao"
	"person_proto/models"
)

var cityDao dao.CityDao = dao.CityDaoImpl{}
var provinceService ProvinceService = ProvinceServiceImpl{}

type CityServiceImpl struct{}

func (CityServiceImpl) CreateCity(data *models.City) error {
	_, err := provinceService.GetProvinceById(data.ProvinceId)
	if err != nil {
		return err
	}
	return cityDao.CreateCity(data)
}
func (CityServiceImpl) GetCityAll() ([]models.City, error) {
	return cityDao.GetCityAll()
}
func (CityServiceImpl) GetCityById(id string) (models.City, error) {
	return cityDao.GetCityById(id)
}

func (CityServiceImpl) UpdateCity(id string, data *models.City) error {
	_, err := cityDao.GetCityById(id)
	if err != nil {
		return err
	}
	return cityDao.UpdateCity(id, data)
}

func (CityServiceImpl) DeleteCity(id string) error {
	_, err := cityDao.GetCityById(id)
	if err != nil {
		return err
	}
	return cityDao.DeleteCity(id)
}
