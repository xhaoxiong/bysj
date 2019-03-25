/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE models
*/
package models

type PageResult struct {
	Code        int         `json:"code"`
	Page        int         `json:"page"` //当前页
	Message     string      `json:"message"`
	Status      int         `json:"status"` //订单状态
	Per         int         `json:"per"`    //每页条数
	Total       int         `json:"total"`  //总数
	UserId      uint        `json:"user_id"`
	OrderNumber string      `json:"order_number"` //订单号
	Search      string      `json:"search"`
	Data        interface{} `json:"data"`
}

type PageCommentResult struct {
	Code    int         `json:"code"`
	HotelId string      `json:"hotel_id"`
	Page    int         `json:"page"` //当前页
	Message string      `json:"message"`
	Status  int         `json:"status"` //评论状态
	Per     int         `json:"per"`    //每页条数
	Total   int         `json:"total"`  //总数
	UserId  uint        `json:"user_id"`
	Search  string      `json:"search"`
	Data    interface{} `json:"data"`
}

type PagePayRecordResult struct {
	Code    int         `json:"code"`
	Page    int         `json:"page"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Per     int         `json:"per"`
	Total   int         `json:"total"`
	Search  string      `json:"search"`
	Data    interface{} `json:"data"`

	UserId    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
	EndAt     string `json:"end_at"`
}
