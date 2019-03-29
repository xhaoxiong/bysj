/**
*@Author: haoxiongxiao
*@Date: 2019/1/28
*@Description: CREATE GO FILE models
 */
package models

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

type Database struct {
	Mysql *gorm.DB
	Redis *redis.Pool
	Mgo   *mgo.Session
}

var DB *Database

/*
*
* @msg mysql的初始化和连接
* @author haoxiong
* @date 2019/1/28 17:27
 */
func openMysqlDB(username, password, addr, name string) *gorm.DB {

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setupDB(db)
	go keepAlive(db)
	return db

}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	//db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(2) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.SingularTable(true)     //设置表名不为负数

	//autoMigrate(db)

}

func (db *Database) Init() {
	DB = &Database{
		Mysql: GetMysqlDB(),
		Redis: GetRedis(),
		Mgo:   GetMgoDB(),
	}
}

func GetMysqlDB() *gorm.DB {
	return InitMysql()
}

func InitMysql() *gorm.DB {
	return openMysqlDB(viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.addr"),
		viper.GetString("mysql.name"))
}

func keepAlive(dbc *gorm.DB) {
	for true {
		dbc.DB().Ping()
		time.Sleep(60 * time.Second)
	}
}

/*
*
* @msg redis 的初始化和连接
* @author haoxiong
* @date 2019/1/28 17:26
 */
func GetRedis() *redis.Pool {
	return InitRedis()
}

func InitRedis() *redis.Pool {
	return openRedisDB(viper.GetString("redis.redis_url"),
		viper.GetInt("redis.redis_idle_timeout_sec"),
		time.Duration(viper.GetInt("redis.redis_idle_timeout_sec")),
		viper.GetString("redis.password"))
}

func openRedisDB(redisURL string, redisMaxIdle int,
	redisIdleTimeoutSec time.Duration, redisPassword string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     redisMaxIdle,
		IdleTimeout: redisIdleTimeoutSec * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisURL)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			//验证redis密码
			//if _, authErr := c.Do("AUTH", redisPassword); authErr != nil {
			//	return nil, fmt.Errorf("redis auth password error: %s", authErr)
			//}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}

/*
*
* @msg mgo连接
* @author haoxiong
* @date 2019/1/29 15:31
 */

func GetMgoDB() *mgo.Session {
	return InitMgo()
}

func InitMgo() *mgo.Session {
	timeout, _ := time.ParseDuration(viper.GetString("timeout"))
	authdb := viper.GetString("authdb")
	authuser := viper.GetString("authuser")
	authpass := viper.GetString("authpass")
	poollimit := viper.GetInt("poollimit")
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{""}, //数据库地址 dbhost: mongodb://user@123456:127.0.0.1:27017
		Timeout:   timeout,      // 连接超时时间 timeout: 60 * time.Second
		Source:    authdb,       // 设置权限的数据库 authdb: admin
		Username:  authuser,     // 设置的用户名 authuser: user
		Password:  authpass,     // 设置的密码 authpass: 123456
		PoolLimit: poollimit,    // 连接池的数量 poollimit: 100
	}
	return openMgoDB(dialInfo)
}

func openMgoDB(dialInfo *mgo.DialInfo) *mgo.Session {
	s, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Errorf(err, "连接mongo 失败")
		panic(err)
	}
	return s
}
