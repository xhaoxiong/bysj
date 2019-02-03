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
	"bysj/services"
	"bysj/models"
)

type AuthController struct {
	Ctx              iris.Context
	WechatApiService wechat_api_services.WechatApiService
	AuthServices     *services.AuthServices
	Common
}

func NewAuthController() *AuthController {
	return &AuthController{AuthServices: services.NewAuthServices()}
}

func (this *AuthController) GetOpenid() {
	code := this.Ctx.FormValue("code")
	userinfo, err := this.WechatApiService.ExchangeUserInfo(code)
	if err != nil {
		this.ReturnJson(10001, "获取用户信息错误")
		return
	}
	this.ReturnSuccess("userinfo", userinfo)
}

func (this *AuthController) PostUserinfo() {
	info := models.UserInfo{}
	if err := this.Ctx.ReadJSON(&info); err != nil {
		log.Error("解析参数错误:", err)
		this.ReturnJson(10001, "解析参数错误")
		return
	}

	if err := this.AuthServices.CreateUser(info); err != nil {
		this.ReturnJson(10001, "添加用户失败")
		return
	}

	this.ReturnSuccess()
}
