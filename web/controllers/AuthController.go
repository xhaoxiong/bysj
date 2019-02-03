/**
*@Author: haoxiongxiao
*@Date: 2019/2/3
*@Description: CREATE GO FILE controller
*/
package controllers

import (
	"bysj/services/wechat_api_services"
	"github.com/kataras/iris"
)

type AuthController struct {
	Ctx              iris.Context
	WechatApiService wechat_api_services.WechatApiService
	Common
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (this *AuthController) PostLogin() {
	code := this.Ctx.FormValue("code")
	userinfo, err := this.WechatApiService.ExchangeUserInfo(code)
	if err != nil {
		this.ReturnJson("获取用户信息错误")
		return
	}

	this.ReturnSuccess("userinfo", userinfo)
}
