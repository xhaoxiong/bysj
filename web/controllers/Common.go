/**
*@Author: haoxiongxiao
*@Date: 2019/2/3
*@Description: CREATE GO FILE controller
*/
package controllers

import "github.com/kataras/iris"

type Common struct {
	Ctx iris.Context
}

func (this *Common) ReturnJson(message string, args ...interface{}) {
	result := make(map[string]interface{})
	result["Status"] = false
	result["Message"] = message

	key := ""

	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)
		default:
			result[key] = arg
		}
	}
	this.Ctx.JSON(result)
	this.Ctx.StopExecution()
	return
}

func (this *Common) ReturnSuccess(args ...interface{}) {
	result := make(map[string]interface{})
	result["Status"] = true
	result["Message"] = "success"
	key := ""
	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)
		default:
			result[key] = arg
		}
	}
	this.Ctx.JSON(result)
	this.Ctx.StopExecution()
	return
}
