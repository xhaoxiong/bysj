/**
*@Author: haoxiongxiao
*@Date: 2019/1/28
*@Description: CREATE GO FILE models
*/
package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	NickName string
	Avatar   string
	Mobile   string `gorm:"unique"`
	Openid   string `gorm:"unique"`
	Gender   string
	Province string
	City     string
	Country  string
	UnionId  string
}
