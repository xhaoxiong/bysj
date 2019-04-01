/**
*@Author: haoxiongxiao
*@Date: 2019/4/1
*@Description: CREATE GO FILE models
 */
package models

import "github.com/jinzhu/gorm"

type City struct {
	gorm.Model
	Val string
}
