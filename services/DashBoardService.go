/**
*@Author: haoxiongxiao
*@Date: 2019/3/25
*@Description: CREATE GO FILE services
*/
package services

import (
	"bysj/models"
	"bysj/repositories"
)

type DashBoardService struct {
	repo *repositories.DashBoardRepositories
}

func NewDashBoardService() *DashBoardService {
	return &DashBoardService{repo: repositories.NewDashBoardRepositories()}
}

//今日订单成交量
func (this *DashBoardService) OrderCount() int {
	return this.repo.OrderCount()
}

//新增用户数
func (this *DashBoardService) UserCount() int {
	return this.repo.UserCount()
}

//今日流水
func (this *DashBoardService) AmountFlow() int {
	return this.repo.AmountFlow()
}

//最近七日订单成交量
func (this *DashBoardService) OrderTrend(orderVolume *[]models.OrderVolume) {
	this.repo.OrderTrend(orderVolume)
}
