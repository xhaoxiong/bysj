/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE models
*/
package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	HotelId string
	UserId  uint
	Content string
}
