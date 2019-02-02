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

	app := iris.New()
	app.Run(iris.Addr(":8080"))
}
