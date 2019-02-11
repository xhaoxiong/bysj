/**
*@Author: haoxiongxiao
*@Date: 2019/2/2
*@Description: CREATE GO FILE sms_api_services
*/
package sms_api_services

import (
	"github.com/xhaoxiong/ShowApiSdk/normalRequest"
	"github.com/garyburd/redigo/redis"
	"bysj/models"
	"bysj/models/redi"
	"github.com/spf13/viper"
)

type SmsApiService struct {
	redi *redis.Pool
}

func NewSmsApiService() *SmsApiService {
	return &SmsApiService{redi: models.DB.Redis}
}

const (
	appId     = 58443
	appSecret = "3bd66623713248bab7c97eabe481bbcd"
)

func (this *SmsApiService) SendSms(code, mobile string) {
	res := normalRequest.ShowapiRequest("http://route.showapi.com/28-1", appId, appSecret)
	res.AddTextPara("mobile", mobile)
	res.AddTextPara("content", "{\"name\":\""+mobile+"\",\"code\":\""+code+"\",\"minute\":\"3\",\"comName\":\"毕设酒店预订平台\"}")
	res.AddTextPara("tNum", "T150606060601")
	res.AddTextPara("big_msg", "")
	_, err := res.Post()

	if err == nil {
		redi.Set(mobile, code)
		redi.SetExpire(mobile, viper.GetString("sms_expire"))
	}
}
