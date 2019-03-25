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
	Prev     int
	After    int
	Amount   int
	Content  string
}
