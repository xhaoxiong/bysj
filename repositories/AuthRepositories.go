/**
*@Author: haoxiongxiao
*@Date: 2019/2/2
*@Description: CREATE GO FILE repositories
 */
package repositories

import (
	"bysj/models"
	"bysj/models/redi"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/core/errors"
	"github.com/lexkong/log"
	"time"
)

type AuthRepositories struct {
	db *gorm.DB
}

func NewAuthRepositories() *AuthRepositories {
	return &AuthRepositories{db: models.DB.Mysql}
}

func (this *AuthRepositories) CreateUser(info *models.UserInfo) (models.User, error) {

	user := models.User{}
	user.Openid = info.Openid
	user.Province = info.Userinfo.Province
	user.Gender = info.Userinfo.Gender
	user.Avatar = info.Userinfo.AvatarUrl
	user.City = info.Userinfo.City
	user.Country = info.Userinfo.Country
	user.NickName = info.Userinfo.NickName
	if this.db.Where("openid = ?", user.Openid).
		First(&models.User{}).RecordNotFound() {
		if err := this.db.Create(&user).Error; err != nil {
			log.Error("添加用户失败", err)
			return user, err
		}
		return user, nil
	} else {

		tx := this.db.Begin()
		u := models.User{}
		this.db.Where("openid = ?", user.Openid).First(&u)
		m := make(map[string]interface{})
		m["id"] = u.ID
		m["openid"] = user.Openid
		m["province"] = user.Province
		m["gender"] = user.Gender
		m["city"] = user.City
		m["country"] = user.Country
		m["nick_name"] = user.NickName
		m["created_at"] = u.CreatedAt
		m["avatar"] = user.Avatar
		m["updated_at"] = time.Now()
		if err := tx.Model(&models.User{}).Updates(m).Error; err != nil {
			return user, err
		}

		if err := tx.Commit().Error; err != nil {
			return user, err
		}

		return user, nil
	}
}

func (this *AuthRepositories) BindUser(mobile, username, cate, cardNum, openid, code string) error {
	user := models.User{}

	if value := redi.GetStringValue(mobile); code != value {
		return errors.New("请输入正确的短信验证码")
	}

	if err := this.db.Where("mobile = ?", mobile).First(&user).Error; err == nil {
		return errors.New("该手机号已经绑定")
	}

	if err := this.db.Where("openid = ?", openid).First(&user).Error; err != nil {
		return errors.New("用户未授权")
	}

	if err := this.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"username":    username,
		"cate":        cate,
		"card_number": cardNum,
		"mobile":      mobile,
		"is_bind":     1,
	}).Error; err != nil {
		return errors.New("绑定失败")
	}
	return nil
}

func (this *AuthRepositories) BindUserCheck(openid string) bool {
	var user models.User
	if err := this.db.Where("openid = ?", openid).First(&user).Error; err != nil {
		return false
	}
	if user.IsBind == 1 {
		return true
	} else {
		return false
	}

}

func (this *AuthRepositories) BindCancel(openid string) bool {
	var user models.User
	if err := this.db.Where("openid = ?", openid).First(&user).Error; err != nil {
		return false
	}

	if err := this.db.Model(&models.User{}).Where("openid = ?", openid).Updates(map[string]interface{}{
		"is_bind": 0,
		"mobile":  "",
	}).Error; err != nil {
		return false
	}

	return true

}

func (this *AuthRepositories) Login() {

}

func (this *AuthRepositories) Logout() {

}

func (this *AuthRepositories) AdminLogin(m map[string]interface{}) (user models.AdminUser, err error) {
	var adminUser models.AdminUser

	if err := this.db.Where("username = ? and password =?",
		m["username"], m["password"]).First(&adminUser).Error; err != nil {
		return adminUser, err
	}
	this.db.Model(&adminUser).Updates(map[string]interface{}{"updated_at": time.Now()})
	return adminUser, nil
}

func (this *AuthRepositories) UserinfoByOpenid(u *models.User) error {
	return this.db.Where("openid = ?", u.Openid).First(&u).Error
}
