/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE models
*/
package models

import "github.com/jinzhu/gorm"

type PayRecord struct {
	gorm.Model
	UserId   uint
	UserName string
	Amount   int
	Content  string
}
