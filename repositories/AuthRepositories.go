/**
*@Author: haoxiongxiao
*@Date: 2019/2/2
*@Description: CREATE GO FILE repositories
 */
package repositories

import "github.com/jinzhu/gorm"

type AuthRepositories struct {
	db *gorm.DB
}

func (this *AuthRepositories) Login() {

}

func (this *AuthRepositories) Logout() {

}
