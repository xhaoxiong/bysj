/**
*@Author: haoxiongxiao
*@Date: 2019/1/28
*@Description: CREATE GO FILE models
*/
package models

import (
	"github.com/lexkong/log"
	"fmt"
	"github.com/spf13/viper"
	"database/sql"
	"github.com/jinzhu/gorm"
)

func autoMigrate(db *gorm.DB) {
	log.Infof("开始同步表")
	if err := db.AutoMigrate(

		&User{}).Error;
		err != nil {
		log.Error("自动建表失败", err)
	}
	log.Infof("同步表成功")
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
		//"Asia/Shanghai"),
		"Local")
	sqlstring := fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8mb4 COLLATE utf8mb4_general_ci", name)
	db, err := sql.Open("mysql", config)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec(sqlstring)

	if err != nil {
		log.Errorf(err, "创建表%s失败", name)
	} else {
		log.Infof("Created Database %s  created", name)
	}
	defer db.Close()
}
