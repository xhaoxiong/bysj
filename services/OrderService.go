/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE services
*/
package services

import (
	"bysj/models"
	"bysj/repositories"
)

type OrderService struct {
	repo *repositories.OrderRepositories
}

func NewOrderService() *OrderService {
	return &OrderService{repo: repositories.NewOrderRepositories()}
}

func (this *OrderService) List(result *models.PageResult) {
	this.repo.List(result)
}

func (this *OrderService) Insert(order *models.Order) error {
	return this.repo.Insert(order)
}

func (this *OrderService) Update(m map[string]interface{}) error {
	return this.repo.Update(m)
}

func (this *OrderService) Delete(ids map[string][]uint) error {
	return this.repo.Delete(ids)
}

func (this *OrderService) NotPayCount(user_id int) int {
	return this.repo.NotPayCount(user_id)
}
