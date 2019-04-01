/**
*@Author: haoxiongxiao
*@Date: 2019/4/1
*@Description: CREATE GO FILE repositories
 */
package repositories

import (
	"bysj/models"
	"github.com/jinzhu/gorm"
)

type HotelRepositories struct {
	db *gorm.DB
}

func NewHotelRepositories() *HotelRepositories {
	return &HotelRepositories{db: models.GetMysqlDB()}
}

func (this *HotelRepositories) CheckCity(cityName string) (string, bool) {
	var city []models.City

	this.db.Where("val like ?", cityName+"%").Find(&city)

	if len(city) > 0 {
		return city[0].Val, true
	}
	return "", false
}
