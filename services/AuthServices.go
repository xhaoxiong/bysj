/**
*@Author: haoxiongxiao
*@Date: 2019/2/3
*@Description: CREATE GO FILE services
*/
package services

import (
	"bysj/repositories"
	"bysj/models"
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
