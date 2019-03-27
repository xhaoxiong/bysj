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
	return &OrderRepositories{db: models.DB.Mysql}
}

func (this *OrderRepositories) List(result *models.PageResult) {

	result.Message = "success"
	result.Code = 10000
	var orders []models.Order

	qs := this.db
	qc := this.db.Model(&models.Order{})

	if result.UserId != 0 {
		qs = qs.Where("user_id = ?", result.UserId)
		qc = qc.Where("user_id = ?", result.UserId)
	}
	if result.Status != 0 {
		qs = qs.Where("status = ?", result.Status)
		qc = qc.Where("status = ?", result.Status)
	}
	s := "%" + result.Search + "%"
	if result.Search != "" {
		qs = qs.Where("order_number like ? or room_info like ? or hotel_item like ?", s, s, s)
		qc = qc.Where("order_number like ? or room_info like ? or hotel_item like ?", s, s, s)
	}

	qs.Limit(result.Per).Preload("User").Offset((result.Page - 1) * result.Per).Find(&orders)
	qc.Count(&result.Total)
	result.Data = orders
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

func (this *OrderRepositories) Delete(ids map[string][]uint) error {
	return this.db.Where("id in (?)", ids["ids"]).Unscoped().Delete(&models.Order{}).Error
}

func (this *OrderRepositories) NotPayCount(user_id int) (count int) {
	this.db.Model(&models.Order{}).Where("status = ? and user_id =?",
		2, user_id).Count(&count)
	return
}
