/**
*@Author: haoxiongxiao
*@Date: 2019/1/27
*@Description: CREATE GO FILE api_services
 */
package hotel_api_services

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/spf13/cast"
	"github.com/xhaoxiong/ShowApiSdk/normalRequest"
)

type roomPriceApiService struct {
	ReqParams RoomPriceReqParams
	ResParams RoomPriceResParams
}

type RoomPriceReqParams struct {
	HotelId string `json:"hotelId"`
	InDate  string `json:"inDate"`
	OutDate string `json:"outDate"`
}

type RoomPriceResParams struct {
	ShowapiResError string `json:"showapi_res_error"`
	ShowapiResID    string `json:"showapi_res_id"`
	ShowapiResCode  int    `json:"showapi_res_code"`
	ShowapiResBody  struct {
		Remark   string `json:"remark"`
		RoomInfo []struct {
			RatePlanInfo []struct {
				InvoiceType    int           `json:"invoiceType"`
				Wifi           string        `json:"wifi"`
				AveragePrice   int           `json:"averagePrice"`
				TaxAndFeePrice int           `json:"taxAndFeePrice"`
				ResNum         int           `json:"resNum"`
				MealInfo       string        `json:"mealInfo"`
				MaxOccupancy   int           `json:"maxOccupancy"`
				Name           string        `json:"name"`
				NetworkInfo    string        `json:"networkInfo"`
				GuestType      int           `json:"guestType"`
				ID             string        `json:"id"`
				PromotionRules []interface{} `json:"promotionRules"`
				Cancel         struct {
					Name string `json:"name"`
					Desc string `json:"desc"`
					Type string `json:"type"`
				} `json:"cancel"`
				BedType        string  `json:"bedType"`
				InstantConfirm bool    `json:"instantConfirm"`
				RatePrice      float64 `json:"ratePrice"`
				BedInfo        string  `json:"bedInfo"`
			} `json:"ratePlanInfo"`
			RoomID     int    `json:"roomId"`
			Floor      string `json:"floor"`
			Facilities []struct {
				TmpSubFacilities []struct {
					SubName string `json:"subName"`
					SubCode string `json:"subCode"`
				} `json:"tmpSubFacilities"`
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"facilities"`
			BedDescription string `json:"bedDescription"`
			Area           string `json:"area"`
			Pictures       []struct {
				Path    string `json:"path"`
				PicName string `json:"picName"`
			} `json:"pictures"`
			RoomNameCn string `json:"roomNameCn"`
			IsExtraBed string `json:"isExtraBed"`
		} `json:"roomInfo"`
		RetCode int `json:"ret_code"`
	} `json:"showapi_res_body"`
}

func (this *roomPriceApiService) GetRoomPrice() (RoomPriceResParams, error) {
	req := normalRequest.ShowapiRequest("http://route.showapi.com/1653-4", appId, appSecret)
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

func ApiRoomPrice(req RoomPriceReqParams) (res RoomPriceResParams, err error) {
	apiService := roomPriceApiService{req, res}

	return apiService.GetRoomPrice()

}
