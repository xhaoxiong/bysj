/**
*@Author: haoxiongxiao
*@Date: 2019/2/3
*@Description: CREATE GO FILE router
*/
package route

import (
	"bysj/web/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func InitRouter(app *iris.Application) {
	mvc.New(app.Party("/auth")).Handle(controllers.NewAuthController())
	mvc.New(app.Party("/hotel")).Handle(controllers.NewHotelController())
	mvc.New(app.Party("/order")).Handle(controllers.NewOrderController())
	mvc.New(app.Party("/comment")).Handle(controllers.NewCommentController())
}
