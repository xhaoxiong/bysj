/**
*@Author: haoxiongxiao
*@Date: 2019/1/26
*@Description: CREATE GO FILE bysj
 */
package main

import (
	"bysj/config"
	"bysj/models"
	"bysj/models/mgodb"
	"bysj/services/hotel_api_services"
	"github.com/kataras/iris"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"

	"bysj/route"
)

var (
	cfg = pflag.StringP("config", "c", "", "./config.yaml")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	models.DB.Init()
	//go services.SyncCity
	if res, err := hotel_api_services.ApiCity(); err != nil {

	} else {
		m := make(map[string]interface{})
		v := res.ShowapiResBody.CityNameList
		m["_id"] = bson.NewObjectId().Hex()
		m["list"] = v
		mgodb.Insert("bysj", "hotel", m)

	}
	app := newApp()
	route.InitRouter(app)
	app.Run(iris.Addr(viper.GetString("addr")))
}

func newApp() *iris.Application {
	app := iris.New()
	//app.Use(middleware.GetJWT().Serve)
	app.Configure(iris.WithOptimizations)
	return app
}
