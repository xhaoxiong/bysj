/**
*@Author: haoxiongxiao
*@Date: 2019/2/11
*@Description: CREATE GO FILE controllers
*/
package controllers

import (
	"bysj/services/hotel_api_services"
	"github.com/kataras/iris"
	"github.com/spf13/cast"
)

type HotelController struct {
	Ctx iris.Context
	Common
}

func NewHotelController() *HotelController {
	return &HotelController{}
}

func (this *HotelController) PostSearch() {

	//reqParams := &hotel_api_services.SearchRequestParams{}

	var reqParams hotel_api_services.SearchRequestParams

	if err := this.Ctx.ReadJSON(&reqParams); err != nil {
		this.ReturnJson(10003, cast.ToString(err))
		return
	}
	res, err := hotel_api_services.ApiSearch(reqParams)
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

func (this *HotelController) PostDetail() {
	var reqParam hotel_api_services.DetailRequestParams
	if err := this.Ctx.ReadJSON(&reqParam); err != nil {
		this.ReturnJson(10002, cast.ToString(err))
		return
	}
	res, err := hotel_api_services.ApiDetail(reqParam)
	if err != nil {
		this.ReturnJson(10001, cast.ToString(err))
		return
	}
	this.ReturnSuccess("data", res.ShowapiResBody.Data)
}

func (this *HotelController) PostRoomPrice() {
	var reqParam hotel_api_services.RoomPriceReqParams

	if err := this.Ctx.ReadJSON(&reqParam); err != nil {
		this.ReturnJson(10001, cast.ToString(err))
		return
	}

	res, err := hotel_api_services.ApiRoomPrice(reqParam)
	if err != nil {
		this.ReturnJson(10002, cast.ToString(err))
		return
	}

	this.ReturnSuccess("data", res.ShowapiResBody.RoomInfo)
}
