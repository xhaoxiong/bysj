/**
*@Author: haoxiongxiao
*@Date: 2019/1/27
*@Description: CREATE GO FILE api_services
*/
package api_services

import (
	"github.com/xhaoxiong/ShowApiSdk/normalRequest"
	"fmt"
	"reflect"
	"github.com/spf13/cast"
	"encoding/json"
	"errors"
)

type RealTimeInquiryService struct {
	ReqParams RealTimeInquiryReqPrams
	ResParams RealTimeInquiryResPrams
}

type RealTimeInquiryReqPrams struct {
	HotelId       string `json:"hotelId"`
	RoomId        string `json:"roomId"`
	NumberOfRooms string `json:"numberOfRooms"`
	InDate        string `json:"inDate"`
	OutDate       string `json:"outDate"`
}

type RealTimeInquiryResPrams struct {
	ShowapiResError string `json:"showapi_res_error"`
	ShowapiResID    string `json:"showapi_res_id"`
	ShowapiResCode  int    `json:"showapi_res_code"`
	ShowapiResBody struct {
		Remark string `json:"remark"`
		Data struct {
			TotalPrice     int    `json:"totalPrice"`
			InvoiceType    int    `json:"invoiceType"`
			Wifi           string `json:"wifi"`
			ID             string `json:"id"`
			AveragePrice   int    `json:"averagePrice"`
			TaxAndFeePrice int    `json:"taxAndFeePrice"`
			NetworkInfo    string `json:"networkInfo"`
			RestNum        int    `json:"restNum"`
			Name           string `json:"name"`
			BedType        string `json:"bedType"`
			MealInfo       string `json:"mealInfo"`
			Cancel struct {
				Name string `json:"name"`
				Desc string `json:"desc"`
				Type string `json:"type"`
			} `json:"cancel"`
			InstantConfirm bool   `json:"instantConfirm"`
			MaxOccupancy   int    `json:"maxOccupancy"`
			RatePrice      int    `json:"ratePrice"`
			BedInfo        string `json:"bedInfo"`
		} `json:"data"`
		RetCode int `json:"ret_code"`
	} `json:"showapi_res_body"`
}

func (this *RealTimeInquiryService) GetRealTimeInquiry() (res RealTimeInquiryResPrams, err error) {
	req := normalRequest.ShowapiRequest("http://route.showapi.com/1653-5", appId, appSecret)
	t := reflect.TypeOf(this.ReqParams)
	v := reflect.ValueOf(this.ReqParams)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			if v.Field(i).CanInterface() {
				key := cast.ToString(t.Field(i).Tag.Get("json"))
				val := cast.ToString(v.Field(i).Interface())

				req.AddTextPara(key, val)
				fmt.Printf("名字:%s    类型:%s  值:%v -标签:%s \n",
					t.Field(i).Name,
					t.Field(i).Type,
					v.Field(i).Interface(),
					t.Field(i).Tag.Get("json"))
			}
		}
	}
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

func ApiRealTimeInquiry(req RealTimeInquiryReqPrams) (res RealTimeInquiryResPrams, err error) {
	apiService := RealTimeInquiryService{req, res}
	return apiService.GetRealTimeInquiry()
}
