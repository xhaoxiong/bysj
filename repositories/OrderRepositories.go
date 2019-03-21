/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE repositories
*/
package repositories

import (
	"bysj/models"
	"github.com/jinzhu/gorm"
)

type OrderRepositories struct {
	db *gorm.DB
}

func NewOrderRepositories() *OrderRepositories {
	return &OrderRepositories{db: models.GetMysqlDB()}
}

func (this *OrderRepositories) List(result *models.PageResult) {

	result.Message = "success"
	result.Code = 10000
	var orders []models.Order

	qs := this.db
	total := 0

	if result.UserId != 0 {
		qs.Where("user_id = ?", result.UserId)
	}
	if result.Status != 0 {
		qs = qs.Where("status = ?", result.Status)
	}
	qs.Limit(result.Count).Offset((result.Page - 1) * result.Count).Find(&orders)
	models.GetMysqlDB().Model(&models.Order{}).Count(&total)
	result.Data = orders
	result.Total = total
}

func (this *OrderRepositories) Insert(order *models.Order) error {

	if err := this.db.Create(&order).Error; err != nil {
		return err
	}
	u := models.User{}

	this.db.Where("id = ?", order.UserId).First(&u)
	order.User = &u
	return nil
}

func (this *OrderRepositories) Update(m map[string]interface{}) error {
	return this.db.Model(&models.Order{}).Updates(m).Error
}

func (this *OrderRepositories) Delete(ids []uint) error {
	return this.db.Where("id in (?)", ids).Unscoped().Delete(&models.Order{}).Error
}
