/**
*@Author: haoxiongxiao
*@Date: 2019/2/11
*@Description: CREATE GO FILE controllers
*/
package controllers

import (
	"github.com/kataras/iris"
	"bysj/services/hotel_api_services"
	"github.com/spf13/cast"
)

type HotelController struct {
	Ctx iris.Context
	Common
}

func NewHotelController() *HotelController {
	return &HotelController{}
}

func (this *HotelController) GetSearch() {

	req := hotel_api_services.SearchRequestParams{
		KeyWord:      "雨花区",
		Page:         "",
		CityName:     "长沙",
		IDate:        "",
		OutDate:      "",
		SortCode:     "",
		ReturnFilter: "1",
		Star:         "",
		Feature:      "",
		MinPrice:     "",
		MaxPrice:     "",
		Facility:     "",
		HotelLabels:  "",
	}
	res, err := hotel_api_services.ApiSearch(req)
	if err != nil {
		this.ReturnJson(10001, cast.ToString(err))
		return
	}

	this.ReturnSuccess("data", res.ShowapiResBody.Data)
}

func (this *HotelController) GetCity() {
	res, err := hotel_api_services.ApiCity()

	if err != nil {
		this.ReturnJson(10001, cast.ToString(err))
		return
	}

	this.ReturnSuccess("data", res.ShowapiResBody.CityNameList)
}
