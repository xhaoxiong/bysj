/**
*@Author: haoxiongxiao
*@Date: 2019/1/26
*@Description: CREATE GO FILE bysj
*/
package main

import (
	"github.com/spf13/pflag"
	"bysj/config"
	"log"
	"github.com/spf13/viper"
	"github.com/kataras/iris"
	"bysj/models"
)

var (
	cfg = pflag.StringP("config", "c", "", "./config.yaml")
)

func init() {

}

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	log.Println(viper.GetString("username"))
	models.DB.Init()
	app := iris.New()
	app.Run(iris.Addr(":8080"))

}
