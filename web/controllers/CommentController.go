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

type CommentController struct {
	Ctx     iris.Context
	Service *services.CommentService
	Common
}

func NewCommentController() *CommentController {
	return &CommentController{Service: services.NewCommentService()}
}

func (this *CommentController) GetList() (result *models.PageCommentResult) {

	if err := this.Ctx.ReadJSON(&result); err != nil {
		this.ReturnJson(10001, cast.ToString(err))
		return
	}

	this.Service.List(result)
	return
}

func (this *CommentController) PostCreate() {
	var comment models.Comment

	if err := this.Ctx.ReadJSON(&comment); err != nil {
		this.ReturnJson(10001, cast.ToString(err))
		return
	}

	if err := this.Service.Create(comment); err != nil {
		this.ReturnJson(10002, cast.ToString(err))
		return
	}

	this.ReturnSuccess()
}
