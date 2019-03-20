/**
*@Author: haoxiongxiao
*@Date: 2019/3/20
*@Description: CREATE GO FILE services
*/
package services

import (
	"bysj/models"
	"bysj/repositories"
)

type UserService struct {
	repo *repositories.UserRepositories
}

func NewUserService() *UserService {
	return &UserService{repo: repositories.NewUserRepositories()}
}

func (this *UserService) List(page, per int, search string) (users []models.User, Total int) {
	return this.repo.List(page, per, search)
}

func (this *UserService) Update(m map[string]interface{}) error {
	return this.repo.Update(m)
}
