/**
*@Author: haoxiongxiao
*@Date: 2019/2/3
*@Description: CREATE GO FILE controller
 */
package controllers

import (
	"bysj/models"
	"bysj/services"
	"bysj/services/sms_api_services"
	"bysj/services/wechat_api_services"
	"bysj/web/middleware"
	"github.com/kataras/iris"
	"github.com/lexkong/log"
	"github.com/spf13/cast"
)

type AuthController struct {
	Ctx              iris.Context
	WechatApiService wechat_api_services.WechatApiService
	AuthServices     *services.AuthServices
	SmsApiService    *sms_api_services.SmsApiService
	Common
}

func NewAuthController() *AuthController {
	return &AuthController{AuthServices: services.NewAuthServices(), SmsApiService: sms_api_services.NewSmsApiService()}
}

func (this *AuthController) PostOpenid() {
	code := this.Ctx.FormValue("code")
	userinfo, err := this.WechatApiService.ExchangeUserInfo(code)
	if err != nil {
		this.ReturnJson(10001, "获取用户信息错误")
		return
	}
	token := middleware.GenerateToken(userinfo.Openid)
	m := make(map[string]interface{})

	m["openid"] = userinfo.Openid
	m["sessionKey"] = userinfo.SessionKey
	m["token"] = token
	this.ReturnSuccess("data", m)
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

//绑定手机号码
func (this *AuthController) PostBind() {
	mobile := this.Ctx.FormValue("mobile")
	username := this.Ctx.FormValue("username")
	cate := this.Ctx.FormValue("cate")
	cardNum := this.Ctx.FormValue("card_number")

	openid := this.Ctx.FormValue("openid")
	code := this.Ctx.FormValue("code")
	//sessionKey := this.Ctx.FormValue("sessionKey")
	if err := this.AuthServices.BindUser(mobile, username, cate, cardNum, openid, code); err != nil {
		this.ReturnJson(10001, cast.ToString(err))
		return
	}
	token := middleware.GenerateToken(openid)

	result := make(map[string]interface{})
	result["message"] = "success"
	result["code"] = 10000
	result["token"] = token
	this.Ctx.JSON(result)
	return
}

func (this *AuthController) PostBindCheck() {
	openid := this.Ctx.FormValue("openid")
	isBind := this.AuthServices.BindUserCheck(openid)
	this.ReturnSuccess("isBind", isBind)
}

func (this *AuthController) PostBindCancel() {
	openid := this.Ctx.FormValue("openid")
	isBind := this.AuthServices.BindCancel(openid)
	this.ReturnSuccess("isBind", isBind)
}

func (this *AuthController) PostGenerateToken() {
	openid := this.Ctx.FormValue("openid")
	token := middleware.GenerateToken(openid)

	m := make(map[string]interface{})
	m["token"] = token
	this.ReturnSuccess("data", m)
	return
}

func (this *AuthController) PostSendSms() {
	mobile := this.Ctx.FormValue("mobile")

	code := this.Krand(6, 0)

	this.SmsApiService.SendSms(code, mobile)

	this.ReturnSuccess()
}
