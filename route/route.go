/**
*@Author: haoxiongxiao
*@Date: 2019/2/3
*@Description: CREATE GO FILE router
*/
package route

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"bysj/web/controllers"
)

func InitRouter(app *iris.Application) {
	mvc.New(app.Party("/auth")).Handle(controllers.NewAuthController())
	mvc.New(app.Party("/hotel")).Handle(controllers.NewHotelController())
}