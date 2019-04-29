/**
*@Author: haoxiongxiao
*@Date: 2019/2/3
*@Description: CREATE GO FILE redis
 */
package redi

import (
	"bysj/models"
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/lexkong/log"
)

func Set(k, v string) {
	c := models.DB.Redis.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v)
	if err != nil {
		log.Error("set error", err)
	}
}

func SetExpire(k, exp string) {
	c := models.DB.Redis.Get()
	defer c.Close()
	_, err := c.Do("expire", k, exp)
	if err != nil {
		log.Error("set error", err)
	}
}

func GetStringValue(k string) string {
	c := models.DB.Redis.Get()
	defer c.Close()
	v, err := redis.String(c.Do("GET", k))
	if err != nil {
		log.Error("Get Error: ", err)
		return ""
	}
	return v
}

func GetZrangeList(k, min, max string) {

}

func GetSliceValue(k string) []string {
	c := models.DB.Redis.Get()
	defer c.Close()
	v, err := redis.Strings(c.Do("GET", k))
	if err != nil {
		log.Error("Get Error: ", err)
		return v
	}
	return v
}

func GetAllKeys(prexMatch string) []string {
	c := models.DB.Redis.Get()
	defer c.Close()
	v, err := redis.Strings(c.Do("keys", prexMatch+"*"))
	if err != nil {
		log.Error("Get Error: ", err)
		return v
	}
	return v
}

func SetKeyExpire(k string, ex int) {
	c := models.DB.Redis.Get()
	defer c.Close()
	_, err := c.Do("EXPIRE", k, ex)
	if err != nil {
		log.Error("set error", err)
	}
}

func CheckKey(k string) bool {
	c := models.DB.Redis.Get()
	defer c.Close()
	exist, err := redis.Bool(c.Do("EXISTS", k))
	if err != nil {
		return false
	} else {
		return exist
	}
}

func DelKey(k string) error {
	c := models.DB.Redis.Get()
	defer c.Close()
	_, err := c.Do("DEL", k)
	if err != nil {
		return err
	}
	return nil
}

func SetJson(k string, data interface{}) error {
	c := models.DB.Redis.Get()
	defer c.Close()
	value, _ := json.Marshal(data)
	n, _ := c.Do("SETNX", k, value)
	if n != int64(1) {
		return errors.New("set failed")
	}
	return nil
}

func getJsonByte(key string) ([]byte, error) {
	c := models.DB.Redis.Get()
	jsonGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return jsonGet, nil
}
