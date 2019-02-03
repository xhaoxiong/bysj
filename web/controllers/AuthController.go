/**
*@Author: haoxiongxiao
*@Date: 2019/2/3
*@Description: CREATE GO FILE controller
*/
package controllers

import (
	"bysj/services/wechat_api_services"
	"github.com/kataras/iris"
	"github.com/lexkong/log"
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
	log.Infof("获取前端的code:%s\n", code)
	userinfo, err := this.WechatApiService.ExchangeUserInfo(code)
	if err != nil {
		this.ReturnJson(10001, "获取用户信息错误")
		return
	}
	log.Infof("获取用户信息:%v", userinfo)
	this.ReturnSuccess("userinfo",userinfo)
}
