/**
*@Author: haoxiongxiao
*@Date: 2019/1/27
*@Description: CREATE GO FILE api_services
 */
package hotel_api_services

import (
	"encoding/json"
	"errors"

	"github.com/xhaoxiong/ShowApiSdk/normalRequest"
)

type DetailApiService struct {
	ReqParams DetailRequestParams
	ResParams DetailResParams
}

type DetailRequestParams struct {
	HotelId string `json:"hotelId"`
}

type DetailResParams struct {
	ShowapiResError string `json:"showapi_res_error"`
	ShowapiResID    string `json:"showapi_res_id"`
	ShowapiResCode  int    `json:"showapi_res_code"`
	ShowapiResBody  struct {
		Data struct {
			EnglishName string  `json:"englishName"`
			Longitude   float64 `json:"longitude"`
			RoomCount   int     `json:"roomCount"`
			Address     string  `json:"address"`
			Pictures    []struct {
				Path string `json:"path"`
				Name string `json:"name"`
			} `json:"pictures"`
			Services []struct {
				Code     string `json:"code"`
				Name     string `json:"name"`
				TypeName string `json:"typeName"`
				TypeCode string `json:"typeCode"`
				Status   string `json:"status"`
			} `json:"services"`
			Tel         string  `json:"tel"`
			Latitude    float64 `json:"latitude"`
			ChineseName string  `json:"chineseName"`
			DebutYear   string  `json:"debutYear"`
			Policy      struct {
				Children          string `json:"children"`
				Pet               string `json:"pet"`
				ArrivalDeparture  string `json:"arrivalDeparture"`
				Requirements      string `json:"requirements"`
				DepositPrepaid    string `json:"depositPrepaid"`
				CheckOutTime      string `json:"checkOutTime"`
				CheckInTime       string `json:"checkInTime"`
				Cancel            string `json:"cancel"`
				AcceptCreditCards string `json:"acceptCreditCards"`
			} `json:"policy"`
			PoiInfos []struct {
				SubPoiInfos []struct {
					Name         string  `json:"name"`
					TrafficeDesc string  `json:"trafficeDesc"`
					Distance     float64 `json:"distance"`
				} `json:"subPoiInfos"`
				Name string `json:"name"`
				Type int    `json:"type"`
			} `json:"poiInfos"`
			Instruction string `json:"instruction"`
			Facilities  []struct {
				Code     string `json:"code"`
				Name     string `json:"name"`
				TypeName string `json:"typeName"`
				TypeCode string `json:"typeCode"`
				Status   string `json:"status"`
			} `json:"facilities"`
			Star         int    `json:"star"`
			DecorateDate string `json:"decorateDate"`
			StarName     string `json:"starName"`
		} `json:"data"`
		Remark     string `json:"remark"`
		UpdateTime string `json:"updateTime"`
		RetCode    int    `json:"ret_code"`
	} `json:"showapi_res_body"`
}

func (this *DetailApiService) GetHotelDetail(hotelId string) (DetailResParams, error) {
	res := normalRequest.ShowapiRequest("http://route.showapi.com/1653-3", appId, appSecret)
	res.AddTextPara("hotelId", hotelId)
	//res.AddFilePara("img", "C:\\Users\\showa\\Desktop\\使用过的\\4.png")//文件上传时设置
	s, err := res.Post()

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

func ApiDetail(req DetailRequestParams) (res DetailResParams, err error) {
	apiService := DetailApiService{req, res}

	return apiService.GetHotelDetail(req.HotelId)
}
