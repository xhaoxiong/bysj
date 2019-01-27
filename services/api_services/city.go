/**
*@Author: haoxiongxiao
*@Date: 2019/1/27
*@Description: CREATE GO FILE api_services
*/
package api_services

import (
	"github.com/xhaoxiong/ShowApiSdk/normalRequest"
	"encoding/json"
	"errors"
)

type CityApiService struct {
	ResParams CityResParams
}

type CityResParams struct {
	ShowapiResError string `json:"showapi_res_error"`
	ShowapiResID    string `json:"showapi_res_id"`
	ShowapiResCode  int    `json:"showapi_res_code"`
	ShowapiResBody struct {
		Remark       string   `json:"remark"`
		RetCode      int      `json:"ret_code"`
		CityNameList []string `json:"cityNameList"`
	} `json:"showapi_res_body"`
}

func (this *CityApiService) GetCities() (res CityResParams, err error) {
	req := normalRequest.ShowapiRequest("http://route.showapi.com/1653-2", appId, appSecret)
	//res.AddFilePara("img", "C:\\Users\\showa\\Desktop\\使用过的\\4.png")//文件上传时设置
	s, err := req.Post()

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(s), &this.ResParams); err != nil {
		panic(err)
	}

	if this.ResParams.ShowapiResCode == 0 {
		return this.ResParams, nil
	} else {
		return this.ResParams, errors.New(this.ResParams.ShowapiResError)
	}
}

func ApiCity() (res CityResParams, err error) {
	apiService := CityApiService{res}

	return apiService.GetCities()
}
