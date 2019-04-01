/**
*@Author: haoxiongxiao
*@Date: 2019/1/26
*@Description: CREATE GO FILE api_services
 */
package hotel_api_services

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/spf13/cast"
	"github.com/xhaoxiong/ShowApiSdk/normalRequest"
)

type SearchApiServices struct {
	ReqParams SearchRequestParams
	Res       SearchRes
}

type SearchRequestParams struct {
	KeyWord      string `json:"keyWords"`     //查询关键字，酒店名称、位置、品牌等
	Page         string `json:"page"`         //页码
	CityName     string `json:"cityName"`     //城市
	IDate        string `json:"iDate"`        //入住时间，格式为：YYYY-MM-DD（默认2天后）
	OutDate      string `json:"outDate"`      //离开时间，格式为：YYYY-MM-DD（默认3天后）
	SortCode     string `json:"sortCode"`     //排序规则(默认1.推荐值排序) 1、推荐值降序 2、起价升序 3、起价降序 6、装修时间排序
	ReturnFilter string `json:"returnFilter"` //是否返回聚合筛选条件,0:否,1:是。注意：returnFilter=1时搜索性能较差，尽量设置returnFilter=0
	Star         string `json:"star"`         //星级 TWO:二星级, THREE:三星级, FOUR:四星级, FIVE:五星级, BUDGET:经济型, CONFORT:舒适型, HIGHEND:高档型, LUXURY:豪华型【多个以逗号:‘,’分隔】
	Feature      string `json:"feature"`      //品牌:通过搜索结果反向聚合
	MinPrice     string `json:"minPrice"`     //房价最低价
	MaxPrice     string `json:"maxPrice"`     //房价最高价
	Facility     string `json:"facility"`     //设施:通过搜索结果反向聚合
	HotelLabels  string `json:"hotellablels"` //特色 1、温泉 3、休闲度假 4、购物便捷 5、客栈民宿 6、青年旅舍 7、精品酒店 8、亲子时刻
}

type SearchRes struct {
	ShowapiResError string `json:"showapi_res_error"`
	ShowapiResID    string `json:"showapi_res_id"`
	ShowapiResCode  int    `json:"showapi_res_code"`
	ShowapiResBody  struct {
		CityName string `json:"cityName"`
		Remark   string `json:"remark"`
		Data     struct {
			HotelList []struct {
				EnglishName string        `json:"englishName"`
				HotelID     int           `json:"hotelId"`
				Longitude   float64       `json:"longitude"`
				Facilities  []interface{} `json:"facilities"`
				Address     string        `json:"address"`
				Latitude    float64       `json:"latitude"`
				Price       int           `json:"price"`
				ChineseName string        `json:"chineseName"`
				Star        int           `json:"star"`
				Picture     string        `json:"picture"`
				StarName    string        `json:"starName"`
			} `json:"hotelList"`
			Count  int `json:"count"`
			Filter []struct {
				FilterName string `json:"filterName"`
				FilterID   string `json:"filterId"`
				Pros       []struct {
					PoiName string `json:"poiName"`
					PoiKey  string `json:"poiKey"`
					Filter  []struct {
						Longitude  float64 `json:"longitude"`
						Code       int     `json:"code"`
						HotelCount int     `json:"hotelCount"`
						Name       string  `json:"name"`
						Heat       int     `json:"heat"`
						Latitude   float64 `json:"latitude"`
					} `json:"filter"`
				} `json:"pros"`
			} `json:"filter"`
		} `json:"data"`
		RetCode int `json:"ret_code"`
	} `json:"showapi_res_body"`
}

func (this *SearchApiServices) GetSearchDataServices() (res SearchRes, err error) {
	req := normalRequest.ShowapiRequest("http://route.showapi.com/1653-1", appId, appSecret)
	t := reflect.TypeOf(this.ReqParams)
	v := reflect.ValueOf(this.ReqParams)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			if v.Field(i).CanInterface() {
				key := cast.ToString(t.Field(i).Tag.Get("json"))
				val := cast.ToString(v.Field(i).Interface())

				req.AddTextPara(key, val)
			}
		}
	}

	s, err := req.Post()
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal([]byte(s), &this.Res); err != nil {
		return res, err
	}

	if this.Res.ShowapiResCode == 0 {
		return this.Res, nil
	} else {
		return this.Res, errors.New(this.Res.ShowapiResError)
	}

}

func ApiSearch(req SearchRequestParams) (res SearchRes, err error) {
	apiService := SearchApiServices{req, res}

	return apiService.GetSearchDataServices()

}
