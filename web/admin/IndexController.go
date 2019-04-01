/**
*@Author: haoxiongxiao
*@Date: 2019/4/1
*@Description: CREATE GO FILE admin
 */
package admin

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type IndexController struct {
	Ctx iris.Context
}

var commandNotLoginIndex = mvc.View{
	Name: "index.html",
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (this *IndexController) Get() mvc.Result {
	return commandNotLoginIndex
}
