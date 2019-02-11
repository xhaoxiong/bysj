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
	"github.com/kataras/iris/core/errors"
	"bysj/models/redi"
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

func (this *AuthRepositories) BindUser(mobile, username, cate, cardNum, openid, code string) error {
	user := models.User{}

	if value := redi.GetStringValue(mobile); code != value {
		return errors.New("请输入正确的验证码")
	}

	if err := this.db.Where("mobile = ?", mobile).First(&user).Error; err == nil {
		return errors.New("该手机号已经注册")
	}

	if err := this.db.Where("openid = ?", openid).First(&user).Error; err != nil {
		return errors.New("用户未授权")
	}

	if err := this.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"username":    username,
		"cate":        cate,
		"card_number": cardNum,
		"mobile":      mobile,
	}).Error; err != nil {
		return errors.New("绑定失败")
	}
	return nil
}

func (this *AuthRepositories) Login() {

}

func (this *AuthRepositories) Logout() {

}
