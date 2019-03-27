/**
*@Author: haoxiongxiao
*@Date: 2019/1/26
*@Description: CREATE GO FILE bysj
 */
package main

import (
	"bysj/config"
	"bysj/models"
	"bysj/route"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	app.Use(crs) //
	app.AllowMethods(iris.MethodOptions)
	//app.Use(middleware.GetJWT().Serve)
	app.Configure(iris.WithOptimizations)
	return app
}
