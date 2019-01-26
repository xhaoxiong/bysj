/**
*@Author: haoxiongxiao
*@Date: 2019/1/26
*@Description: CREATE GO FILE api_services
*/
package api_services

type CommonRequestParams struct {
	ShowapiAppid     string `json:"showapi_appid"`     //100
	ShowapiSign      string `json:"showapi_sign"`      //698d51a19d8a121ce581499d7b701668
	ShowapiResGzip   string `json:"showapi_res_gzip"`  //1或0
	ShowapiTimestamp string `json:"showapi_timestamp"` //20141114142239
}

/*
ShowapiResCode
易源返回标志,0为成功，其他为失败。
0成功
-1，系统调用错误
-2，可调用次数或金额为0
-3，读取超时
-4，服务端返回数据解析错误
-5，后端服务器DNS解析错误
-6，服务不存在或未上线
-1000，系统维护
-1002，showapi_appid字段必传
-1003，showapi_sign字段必传
-1004，签名sign验证有误
-1005，showapi_timestamp无效
-1006，app无权限调用接口
-1007，没有订购套餐
-1008，服务商关闭对您的调用权限
-1009，调用频率受限
-1010，找不到您的应用
-1011，子授权app_child_id无效
-1012，子授权已过期或失效
-1013，子授权ip受限
*/

type CommonResponseParams struct {
	ShowapiResBody  string `json:"showapi_res_body"` //消息体的JSON封装，所有应用级的返回参数将嵌入此对象 。
	ShowapiResCode  string `json:"showapi_res_code"`
	ShowapiResError string `json:"showapi_res_error"` //错误信息的展示
	ShowapiResId    string `json:"showapi_res_id"`    //本次请求id
}
