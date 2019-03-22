/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE models
*/
package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	HotelId   string
	HotelName string
	UserId    uint
	User      *User
	Content   string
	OrderId   uint
	Order     *Order
	Status    int //评论状态
}
