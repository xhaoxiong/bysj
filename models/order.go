/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE models
*/
package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	OrderNumber string
	HotelId     string
	RoomId      string
	UserId      uint
	Status      int //1 预下单 2待支付 3已支付 4已取消 5待评价
	RoomInfo    interface{}
}
