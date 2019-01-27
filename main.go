/**
*@Author: haoxiongxiao
*@Date: 2019/1/26
*@Description: CREATE GO FILE bysj
*/
package main

import (
	"bysj/services/api_services"
	"log"
)

func main() {

	//req := api_services.SearchRequestParams{
	//	KeyWord:      "",
	//	Page:         "",
	//	CityName:     "长沙",
	//	IDate:        "",
	//	OutDate:      "",
	//	SortCode:     "",
	//	ReturnFilter: "",
	//	Star:         "",
	//	Feature:      "
	//	MinPrice:     "",
	//	MaxPrice:     "",
	//	Facility:     "",
	//	HotelLabels:  "",
	//}
	//log.Println(req)
	//req2 := api_services.RoomPriceReqParams{
	//	HotelId: "540562",
	//}

	//log.Println(api_services.ApiSearch(req))
	//log.Println(api_services.ApiRoomPrice(req2))

	log.Println(api_services.ApiCity())
}
