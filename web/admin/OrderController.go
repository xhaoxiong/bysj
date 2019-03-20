/**
*@Author: haoxiongxiao
*@Date: 2019/3/20
*@Description: CREATE GO FILE admin
*/
package admin

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
		return
	}
	this.Service.List(result)
	return
}

func (this *OrderController) PostUpdate() {
	m := make(map[string]interface{})
	if err := this.Ctx.ReadJSON(&m); err != nil {
		this.ReturnJson(10001, cast.ToString(err))
		return
	}

	if err := this.Service.Update(m); err != nil {
		this.ReturnJson(10002, "更新失败")
		return
	}
	this.ReturnSuccess()
}

func (this *OrderController) PostDelete() {
	m := make(map[string]interface{})

	if err := this.Ctx.ReadJSON(m); err != nil {
		this.ReturnJson(10001, cast.ToString(err))
		return
	}

	if err := this.Service.Delete(m["ids"].([]uint)); err != nil {
		this.ReturnJson(10002, cast.ToString(err))
		return
	}
	this.ReturnSuccess()
}
