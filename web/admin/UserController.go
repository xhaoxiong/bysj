/**
*@Author: haoxiongxiao
*@Date: 2019/3/20
*@Description: CREATE GO FILE admin
*/
package admin

import (
	"bysj/services"
	"github.com/kataras/iris"
	"github.com/spf13/cast"
)

type UserController struct {
	Ctx iris.Context
	Common
	Service *services.UserService
}

func NewUserController() *UserController {
	return &UserController{Service: services.NewUserService()}
}

func (this *UserController) GetList() {
	page := cast.ToInt(this.Ctx.FormValue("page"))
	per := cast.ToInt(this.Ctx.FormValue("per"))
	search := this.Ctx.FormValue("search")

	users, Total := this.Service.List(page, per, search)

	this.ReturnSuccess("data", users, "total", Total)
}

func (this *UserController) PostUpdate() {
	m := make(map[string]interface{})
	if err := this.Ctx.ReadJSON(&m); err != nil {
		this.ReturnJson(10001, cast.ToString(err))
		return
	}
	if err := this.Service.Update(m); err != nil {
		this.ReturnJson(10002, cast.ToString(err))
		return
	}
	this.ReturnSuccess()
}
