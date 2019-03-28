/**
*@Author: haoxiongxiao
*@Date: 2019/3/25
*@Description: CREATE GO FILE admin
 */
package admin

import (
	"bysj/models"
	"bysj/services"
	"github.com/kataras/iris"
)

type DashBoardController struct {
	Ctx     iris.Context
	Service *services.DashBoardService
	Common
}

func NewDashBoardController() *DashBoardController {
	return &DashBoardController{Service: services.NewDashBoardService()}
}

//今日订单成交量
func (this *DashBoardController) PostOrderCount() {
	count := this.Service.OrderCount()
	this.ReturnSuccess("data", count)
}

//新增用户数
func (this *DashBoardController) PostUserCount() {
	count := this.Service.UserCount()
	this.ReturnSuccess("data", count)
}

//今日流水
func (this *DashBoardController) PostAmountFlow() {
	flow := this.Service.AmountFlow()
	this.ReturnSuccess("data", flow)
}

//最近七日订单成交量
func (this *DashBoardController) PostOrderTrend() {
	var orderVolume []models.OrderVolume
	this.Service.OrderTrend(&orderVolume)

	this.ReturnSuccess("data", orderVolume)
}
