/**
*@Author: haoxiongxiao
*@Date: 2019/3/28
*@Description: CREATE GO FILE models
 */
package models

import "github.com/jinzhu/gorm"

type FeedBack struct {
	gorm.Model
	User    *User
	UserId  uint
	Content string `gorm:"text"`
}

