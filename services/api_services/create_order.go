/**
*@Author: haoxiongxiao
*@Date: 2019/1/27
*@Description: CREATE GO FILE api_services
*/
package api_services

type CreateOrderApiService struct{}

type CreateOrderReqParams struct {
	CustomerName   string `json:"customerName"`
	RatePlanId     string `json:"ratePlanId"`
	HotelId        string `json:"hotelId"`
	SpecialRemarks string `json:"specialRemarks"`
	ContactPhone   string `json:"contactPhone"`
	ContactEmail   string `json:"contactEmail"`
	InDate         string `json:"inDate"`
	OutDate        string `json:"outDate"`
}


