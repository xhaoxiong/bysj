/**
*@Author: haoxiongxiao
*@Date: 2019/3/28
*@Description: CREATE GO FILE admin
 */
package admin

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
	if feedback.Content == "" {
		this.ReturnJson(10001, "反馈失败")
		return
	}
	if err := this.Service.Create(&feedback); err != nil {
		this.ReturnJson(10001, "反馈失败")
		return
	}
	this.ReturnSuccess()
}

func (this *FeedBackController) PostList() (result *models.PageFeedBackResult) {
	if err := this.Ctx.ReadJSON(&result); err != nil {
		this.ReturnJson(10001, "获取列表失败")
		return
	}

	this.Service.List(result)
	return
}
