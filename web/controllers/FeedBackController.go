/**
*@Author: haoxiongxiao
*@Date: 2019/3/28
*@Description: CREATE GO FILE controllers
 */
package controllers

import (
	"bysj/models"
	"bysj/services"
	"github.com/kataras/iris"
)

type FeedBackController struct {
	Ctx     iris.Context
	Service *services.FeedBackService
	Common
}

func NewFeedBackController() *FeedBackController {
	return &FeedBackController{Service: services.NewFeedBackService()}
}

func (this *FeedBackController) PostCreate() {
	var feedback models.FeedBack
	if err := this.Ctx.ReadJSON(&feedback); err != nil {
		this.ReturnJson(10001, "反馈失败")
		return
	}

	if err := this.Service.Create(&feedback); err != nil {
		this.ReturnJson(10001, "反馈失败")
		return
	}
	this.ReturnSuccess()
}
