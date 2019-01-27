/**
*@Author: haoxiongxiao
*@Date: 2019/1/26
*@Description: CREATE GO FILE api_services
*/
package api_services

import (
	"testing"
	"log"
)

func TestNewSearchApiServices(t *testing.T) {
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

}
