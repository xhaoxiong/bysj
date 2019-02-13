/**
*@Author: haoxiongxiao
*@Date: 2019/1/28
*@Description: CREATE GO FILE models
*/
package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	NickName string
	Avatar   string
	Mobile   string
	Openid   string `gorm:"unique"`
	Gender   int
	Province string
	City     string
	Country  string
	UnionId  string

	Username   string
	CardNumber string
	Cate       int //1 2 3 4

	IsBind int `gorm:"default:2"` //1绑定 2未绑定
}

type UserInfo struct {
	Id int `json:"id"`
	Userinfo struct {
		AvatarUrl string `json:"avatarUrl"`
		City      string `json:"city"`
		Country   string `json:"country"`
		Gender    int    `json:"gender"`
		Language  string `json:"language"`
		NickName  string `json:"nickName"`
		Province  string `json:"province"`
	} `json:"userinfo"`
	Openid string `json:"openid"`
}
