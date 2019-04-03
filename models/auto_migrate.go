/**
*@Author: haoxiongxiao
*@Date: 2019/1/28
*@Description: CREATE GO FILE models
 */
package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
)

func autoMigrate(db *gorm.DB) {

	if err := db.AutoMigrate(
		&AdminUser{},
		&User{},
		&Order{},
		&PayRecord{},
		&Comment{},
		&FeedBack{},
		&OrderVolume{},
		&City{},
	).Error;
		err != nil {
		log.Error("自动建表失败", err)
	}
}