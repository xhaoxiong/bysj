/**
*@Author: haoxiongxiao
*@Date: 2019/4/1
*@Description: CREATE GO FILE services
 */
package services

import "bysj/repositories"

type HotelService struct {
	repo *repositories.HotelRepositories
}

func NewHotelService() *HotelService {
	return &HotelService{repo: repositories.NewHotelRepositories()}
}

func (this *HotelService) CheckCity(cityName string) bool {
	return this.repo.CheckCity(cityName)
}
