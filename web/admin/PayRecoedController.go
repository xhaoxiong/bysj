/**
*@Author: haoxiongxiao
*@Date: 2019/3/25
*@Description: CREATE GO FILE admin
*/
package admin

import (
	"bysj/models"
	"bysj/services"
	"fmt"
	"github.com/kataras/iris"
)

type PayRecordController struct {
	Ctx     iris.Context
	Service *services.PayRecordService
	Common
}

func NewPayRecordController() *PayRecordController {
	return &PayRecordController{Service: services.NewPayRecordService()}
}

func (this *PayRecordController) PostList() (result *models.PagePayRecordResult) {

	if err := this.Ctx.ReadJSON(&result); err != nil {
		fmt.Println(err)
		this.ReturnJson(10001, "解析参数错误")
		return
	}
	this.Service.List(result)
	return
}

func (this *PayRecordController) PostCreate() {
	var payRecord *models.PayRecord
	if err := this.Ctx.ReadJSON(&payRecord); err != nil {
		fmt.Println(err)
		this.ReturnJson(10001, "解析参数错误")
		return
	}

	if err := this.Service.Create(payRecord); err != nil {
		fmt.Println(err)
		this.ReturnJson(10002, "创建失败")
		return
	}

	this.ReturnSuccess()

}
