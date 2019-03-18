/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE controllers
*/
package controllers

import (
	"bysj/models"
	"bysj/services"
	"github.com/kataras/iris"
	"github.com/spf13/cast"
)

type OrderController struct {
	Ctx     iris.Context
	Service *services.OrderService
	Common
}

func NewOrderController() *OrderController {
	return &OrderController{Service: services.NewOrderService()}
}

func (this *OrderController) GetList() (result *models.PageResult) {
	if err := this.Ctx.ReadJSON(&result); err != nil {
		this.ReturnJson(10001, cast.ToString(err))
		return
	}
	return
}

func (this *OrderController) PostCreate() {
	var order *models.Order

	if err := this.Ctx.ReadJSON(&order); err != nil {
		this.ReturnJson(10001, cast.ToString(err))
		return
	}
	if err := this.Service.Insert(order); err != nil {
		this.ReturnJson(10002, cast.ToString(err))
		return
	}
	this.ReturnSuccess()
}

func (this *OrderController) PostUpdate() {
	m := make(map[string]interface{})
	if err := this.Ctx.ReadJSON(m); err != nil {
		this.ReturnJson(10001, cast.ToString(err))
		return
	}

	if err := this.Service.Update(m); err != nil {
		this.ReturnJson(10002, cast.ToString(err))
	}
	this.ReturnSuccess()
}


