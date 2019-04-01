/**
*@Author: haoxiongxiao
*@Date: 2019/1/28
*@Description: CREATE GO FILE models
 */
package models

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

func autoMigrate(db *gorm.DB) {

	if err := db.AutoMigrate(
		&AdminUser{},
		&User{},
		&Order{},
		&PayRecord{},
		&Comment{},
		&FeedBack{},
		&OrderVolume{},
		&City{},
	).Error;
		err != nil {
		log.Error("自动建表失败", err)
	}
}

func Syncdb() {
	//autoMigrate()
	CreateDB()
	InitMysql()

}

func CreateDB() {
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	addr := viper.GetString("mysql.addr")
	name := viper.GetString("mysql.name")

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Asia/Shanghai")
	//"Local")
	sqlstring := fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8mb4 COLLATE utf8mb4_general_ci", name)
	db, err := sql.Open("mysql", config)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec(sqlstring)

	if err != nil {
		log.Errorf(err, "创建表%s失败", name)
	} else {
		fmt.Println("Created Database %s  created", name)
	}
	defer db.Close()
}
