/**
*@Author: haoxiongxiao
*@Date: 2019/2/2
*@Description: CREATE GO FILE repositories
 */
package repositories

import (
	"github.com/jinzhu/gorm"
	"bysj/models"
	"github.com/lexkong/log"
)

type AuthRepositories struct {
	db *gorm.DB
}

func NewAuthRepositories() *AuthRepositories {
	return &AuthRepositories{db: models.DB.Mysql}
}

func (this *AuthRepositories) CreateUser(info models.UserInfo) error {

	user := models.User{}
	user.Openid = info.Openid
	user.Province = info.Userinfo.Province
	user.Gender = info.Userinfo.Gender
	user.Avatar = info.Userinfo.AvatarUrl
	user.City = info.Userinfo.City
	user.Country = info.Userinfo.Country

	if err := this.db.Create(&user).Error; err != nil {
		log.Error("添加用户失败", err)
		return err
	}
	return nil
}

func (this *AuthRepositories) Login() {

}

func (this *AuthRepositories) Logout() {

}
