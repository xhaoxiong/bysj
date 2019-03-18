/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE main
 */
package main

import (
	"bysj/services/hotel_api_services"
	"encoding/json"
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	S := `{"keywords":"","page":"","cityName":"长沙","iDate":"","outDate":"","sortCode":"","returnFilter":"","star":"","feature":"","minPrice":"","maxPrice":"","facility":"","hotellablels":""}`

	var params hotel_api_services.SearchRequestParams
	json.Unmarshal([]byte(S), &params)
	fmt.Println(params.Page)
	fmt.Println(params.CityName)
	fmt.Println(params.KeyWord)
}
