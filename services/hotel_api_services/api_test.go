/**
*@Author: haoxiongxiao
*@Date: 2019/1/26
*@Description: CREATE GO FILE api_services
 */
package hotel_api_services

import (
	"testing"

	"bysj/services/sms_api_services"
)

func TestNewSearchApiServices(t *testing.T) {

	a := &sms_api_services.SmsApiService{}
	a.SendSms("123456", "18374878791")
	/*
		req := SearchRequestParams{
			KeyWord:      "keyWord",
			Page:         "page",
			CityName:     "cityName",
			IDate:        "iDate",
			OutDate:      "outDate",
			SortCode:     "sortCode",
			ReturnFilter: "returnFilter",
			Star:         "star",
			Feature:      "feature",
			MinPrice:     "minPrice",
			MaxPrice:     "maxPrice",
			Facility:     "facility",
			HotelLabels:  "hotelLabels",
		}
		log.Println(ApiSearch(req))


	/*
		res := normalRequest.ShowapiRequest("http://route.showapi.com/1653-3", appId, appSecret)
		res.AddTextPara("hotelId", "474138")
		//res.AddFilePara("img", "C:\\Users\\showa\\Desktop\\使用过的\\4.png")//文件上传时设置
		s, err := res.Post()

		if err != nil {
			panic(err)
		}

		resParams := DetailResParams{}

		if err := json.Unmarshal([]byte(s), &resParams); err != nil {
			panic(err)
		}

		log.Println(resParams)

	*/

	/*
		res := normalRequest.ShowapiRequest("http://route.showapi.com/1653-4", appId, appSecret)
		res.AddTextPara("hotelId", "540562")
		res.AddTextPara("inDate", "")
		res.AddTextPara("outDate", "")
		//res.AddFilePara("img", "C:\\Users\\showa\\Desktop\\使用过的\\4.png")//文件上传时设置
		fmt.Println(res.Post())

		/*
		res := normalRequest.ShowapiRequest("http://route.showapi.com/1653-2", appId, appSecret)
		//res.AddFilePara("img", "C:\\Users\\showa\\Desktop\\使用过的\\4.png")//文件上传时设置
		fmt.Println(res.Post())
	*/

	/*
		res := normalRequest.ShowapiRequest("http://route.showapi.com/1653-5", appId, appSecret)
		res.AddTextPara("hotelId", "540562")
		res.AddTextPara("roomId", "1_540562_200_222098405_31029_3_113052_2_12_11198_8")
		res.AddTextPara("numberOfRooms", "1")
		res.AddTextPara("inDate", "2019-01-28")
		res.AddTextPara("outDate", "2019-01-29")
		//res.AddFilePara("img", "C:\\Users\\showa\\Desktop\\使用过的\\4.png")//文件上传时设置
		fmt.Println(res.Post())
	*/
	/*
	res := normalRequest.ShowapiRequest("http://route.showapi.com/1653-6", appId, appSecret)
	res.AddTextPara("showapi_timestamp", "20190127230334")
	res.AddTextPara("customerName", "xxx")
	res.AddTextPara("ratePlanId", "1_2348800_1502_221845747_23690_3_17857489")
	res.AddTextPara("hotelId", "2348800")
	res.AddTextPara("specialRemarks", "0")
	res.AddTextPara("contactName", "xxx")
	res.AddTextPara("contactPhone", "18374878792")
	res.AddTextPara("contactEmail", "")
	res.AddTextPara("inDate", "2019-01-28")
	res.AddTextPara("outDate", "2019-01-29")
	//res.AddFilePara("img", "C:\\Users\\showa\\Desktop\\使用过的\\4.png")//文件上传时设置
	fmt.Println(res.Post())
	*/
}
