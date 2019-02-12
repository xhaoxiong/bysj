/**
*@Author: haoxiongxiao
*@Date: 2019/1/26
*@Description: CREATE GO FILE bysj
 */
package main

import (
	"bysj/config"
	"bysj/models"

	"github.com/kataras/iris"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"bysj/route"
	"bysj/web/middleware"
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

	app := newApp()
	route.InitRouter(app)
	app.Run(iris.Addr(viper.GetString("addr")))
}

func newApp() *iris.Application {
	app := iris.New()
	app.Use(middleware.GetJWT().Serve)
	app.Configure(iris.WithOptimizations)
	return app
}
