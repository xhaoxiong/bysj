/**
*@Author: haoxiongxiao
*@Date: 2019/2/3
*@Description: CREATE GO FILE wechat_api_services
*/
package wechat_api_services

import (
	"net/http"
	"github.com/spf13/viper"
	"io/ioutil"
)

type WechatApiService struct{}

func (this *WechatApiService) ExchangeUserInfo(code string) (userinfo interface{}, err error) {
	appId := viper.GetString("mini_program.app_id")
	appSecret := viper.GetString("mini.program.app_secret")

	resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?" +
		"appid=" + appId + "&" +
		"secret=" + appSecret + "&" +
		"js_code=" + code + "&" +
		"grant_type=authorization_code")

	if err != nil {
		return nil, err
	}

	bytes, e := ioutil.ReadAll(resp.Body)

	return string(bytes), e
}
