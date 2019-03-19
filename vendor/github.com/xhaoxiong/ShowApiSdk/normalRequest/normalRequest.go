package normalRequest

import (
	"strings"
	"net/url"
	"time"
	"github.com/xhaoxiong/ShowApiSdk/httplib"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"encoding/base64"
	"errors"
)

type NormalReq struct {
	url            string
	badyParas      url.Values
	file           map[string]string
	readTimeOut    time.Duration
	connectTimeOut time.Duration
	headParas      url.Values
}

//用于请求官网
func ShowapiRequest(reqUrl string, appid int, sign string) *NormalReq {
	values := make(url.Values)
	values.Set("showapi_appid", strconv.Itoa(appid))
	values.Set("showapi_sign", sign)
	return &NormalReq{reqUrl, values, make(map[string]string), 3 * time.Second, 15 * time.Second, make(url.Values)}
}

//通用请求
func NormalRequest(reqUrl string) *NormalReq {
	values := make(url.Values)
	return &NormalReq{reqUrl, values, make(map[string]string), 3 * time.Second, 15 * time.Second, make(url.Values)}
}

func (request *NormalReq) AddTextPara(key, value string) {
	request.badyParas.Set(key, value)
}

func (request *NormalReq) AddFilePara(key, fileName string) {
	request.file[key] = fileName
}
func (request *NormalReq) AddHeadPara(key, value string) {
	request.headParas.Set(key, value)
}
func (request *NormalReq) SetReadTimeOut(readTimeOut time.Duration) {
	request.readTimeOut = readTimeOut
}
func (request *NormalReq) SetConnectTimeOut(connectTimeOut time.Duration) {
	request.connectTimeOut = connectTimeOut
}

//get请求
func (request *NormalReq) Get() (string, error) {
	req := httplib.Get(strings.TrimSpace(request.url) + "?" + request.badyParas.Encode())
	req.SetTimeout(request.connectTimeOut, request.readTimeOut)
	for k, v := range request.headParas {
		req.Header(k, v[0])
	}
	return req.String()
}

//post请求包括文件上传部分
func (request *NormalReq) Post() (string, error) {
	req := httplib.Post(strings.TrimSpace(request.url))
	for k, v := range request.badyParas {
		req.Param(k, v[0])
	}
	for k, v := range request.file {
		req.PostFile(k, v)
	}
	req.SetTimeout(request.connectTimeOut, request.readTimeOut)
	for k, v := range request.headParas {
		req.Header(k, v[0])
	}
	return req.String()
}

//获取返回头

//解析json
func ParseJson(req string) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(req), &data); err != nil {
		return nil, errors.New("showapi return body is nil")
	}
	return data, nil
}

//图片文件传base64
func Base64(fileName string) (string) {
	fileBase64, _ := ioutil.ReadFile(fileName)
	return base64.StdEncoding.EncodeToString(fileBase64)
}
