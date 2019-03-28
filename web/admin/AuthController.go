/**
*@Author: haoxiongxiao
*@Date: 2019/3/20
*@Description: CREATE GO FILE admin
 */
package admin

import (
	"bysj/services"
	"bysj/web/middleware"
	"github.com/kataras/iris"
	"github.com/spf13/cast"
)

type AuthController struct {
	Ctx     iris.Context
	Service *services.AuthServices
	Common
}

func NewAuthController() *AuthController {
	return &AuthController{Service: services.NewAuthServices()}
}

func (this *AuthController) PostLogin() {
	m := make(map[string]interface{})
	if err := this.Ctx.ReadJSON(&m); err != nil {
		this.ReturnJson(10001, cast.ToString(err))
		return
	}

	if user, err := this.Service.AdminLogin(m); err != nil {
		this.ReturnJson(10002, cast.ToString(err))
		return
	} else {

		token := middleware.GenrateAdminToken(&user)
		result := make(map[string]interface{})
		result["code"] = 10000
		result["message"] = "success"
		result["token"] = token
		result["userinfo"] = user
		this.Ctx.JSON(result)
		return
	}
}
