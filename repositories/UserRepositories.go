/**
*@Author: haoxiongxiao
*@Date: 2019/2/2
*@Description: CREATE GO FILE repositories
 */
package repositories

import (
	"bysj/models"
	"github.com/jinzhu/gorm"
)

type UserRepositories struct {
	db *gorm.DB
}

func NewUserRepositories() *UserRepositories {
	return &UserRepositories{db: models.GetMysqlDB()}
}

func (this *UserRepositories) List(page, per int, search string) (users []models.User, Total int) {

	qs := this.db
	qc := this.db.Model(&models.User{})
	if per == 0 {
		per = 10
	}

	s := "%" + search + "%"
	if search != "" {
		qs = qs.Where("nick_name like ? or username like ? ", s, s)
		qc = qc.Where("nick_name like ? or username like ? ", s, s)
	}
	qc.Count(&Total)
	qs.Limit(per).Offset((page - 1) * per).Find(&users)
	return
}

func (this *UserRepositories) Update(m map[string]interface{}) error {
	return this.db.Model(&models.User{}).Updates(m).Error
}
