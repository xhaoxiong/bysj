/**
*@Author: haoxiongxiao
*@Date: 2019/2/3
*@Description: CREATE GO FILE services
*/
package services

import (
	"bysj/models"
	"bysj/repositories"
)

type AuthServices struct {
	repo *repositories.AuthRepositories
}

func NewAuthServices() *AuthServices {
	return &AuthServices{repo: repositories.NewAuthRepositories()}
}

func (this *AuthServices) CreateUser(info models.UserInfo) error {
	return this.repo.CreateUser(info)
}

func (this *AuthServices) BindUser(mobile, username, cate, cardNum, openid, code string) error {
	return this.repo.BindUser(mobile, username, cate, cardNum, openid, code)
}

func (this *AuthServices) BindUserCheck(openid string) bool {
	return this.repo.BindUserCheck(openid)
}

func (this *AuthServices) BindCancel(openid string) bool {
	return this.repo.BindCancel(openid)
}

func (this *AuthServices) AdminLogin(m map[string]interface{}) (user models.AdminUser, err error) {
	return this.repo.AdminLogin(m)
}
