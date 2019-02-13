/**
*@Author: haoxiongxiao
*@Date: 2019/2/13
*@Description: CREATE GO FILE models
*/
package main

import (
	"testing"
	"github.com/spf13/pflag"
	"bysj/config"
	"bysj/models/mgodb"
	"time"
	"gopkg.in/mgo.v2/bson"
	"bysj/models"
	"fmt"
)

func TestInitMgo(t *testing.T) {
	type Test struct {
		Id      string
		Title   string
		Content string
		Date    time.Time
	}
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	models.DB.Init()

	test := Test{
		Id:      bson.NewObjectId().Hex(),
		Title:   "标题",
		Content: "内容",
		Date:    time.Now(),
	}
	mgodb.Insert("test", "testModel", test)

	//test 查询title="标题",并且返回结果中去除`_id`字段
	var result []Test
	mgodb.FindAll("test", "testModel", bson.M{"title": "标题"}, bson.M{"_id": 0}, &result)
	fmt.Println(result)

}
