/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE models
*/
package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model

	Amount      int
	OrderNumber string
	RealId      string //实际居住的某一间id
	RealInfo    string //实际的房间的具体信息
	HotelId     string //酒店id
	HotelItem   string
	RoomId      string //房间id
	RoomInfo    string
	User        *User
	UserId      uint
	Status      int //1 预下单 2待支付 3已支付 4已取消 5待评价
	InDate      string
	OutDate     string
}

//每天的订单成交量
type OrderVolume struct {
	Date   string
	Volume int
}
