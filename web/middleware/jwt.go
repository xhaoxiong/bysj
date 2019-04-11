/**
*@Author: haoxiongxiao
*@Date: 2019/2/12
*@Description: CREATE GO FILE middleware
 */
package middleware

import (
	"bysj/models"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"

	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"strings"
	"time"
)

var JwtAuthMiddleware = jwtmiddleware.New(jwtmiddleware.Config{
	ValidationKeyGetter: validationKeyGetterFuc,
	SigningMethod:       jwt.SigningMethodHS256,
	Expiration:          true,
	Extractor:           extractor,
}).Serve

const jwtKey = "bysj"

var validationKeyGetterFuc = func(token *jwt.Token) (interface{}, error) {
	return []byte(jwtKey), nil
}

var extractor = func(ctx context.Context) (string, error) {
	authHeader := ctx.GetHeader("token")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	return authHeader, nil
}

//注册jwt中间件
func GetJWT() *jwtmiddleware.Middleware {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(jwtKey), nil
		},
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		//ErrorHandler: func(context.Context, string)
		ErrorHandler: func(ctx iris.Context, s string) {
			if strings.Contains(ctx.Request().RequestURI, "/auth/openid") ||
				strings.Contains(ctx.Request().RequestURI, "/auth/bind") ||
				strings.Contains(ctx.Request().RequestURI, "/auth/send/sms") ||
				strings.Contains(ctx.Request().RequestURI, "/auth/userinfo") ||
				strings.Contains(ctx.Request().RequestURI, "/auth/generate/token") ||
				strings.Contains(ctx.Request().RequestURI, "/api/admin/auth/login") ||
				strings.Contains(ctx.Request().RequestURI, "/hotel/city/check") ||
				strings.Contains(ctx.Request().RequestURI, "/admin") ||
				strings.Contains(ctx.Request().RequestURI, "/hotel/search") {
				ctx.Next()
			} else {
				result := make(map[string]interface{})
				result["msg"] = "认证失败"
				result["code"] = 10014
				ctx.JSON(result)
			}

		},
	})
	return jwtHandler
}

func GenrateAdminToken(user *models.AdminUser) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user, //openid
		//"sessionKey": sessionKey,                                               //sessionKey
		"iss": "iris_bysj",                                              //签发者
		"iat": time.Now().Unix(),                                        //签发时间
		"jti": "9527",                                                   //jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
		"exp": time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(), //过期时间
	})
	tokenString, _ := token.SignedString([]byte(jwtKey))
	fmt.Println("签发时间：", time.Now().Unix())
	fmt.Println("到期时间：", time.Now().Add(10 * time.Hour * time.Duration(1)).Unix())
	return tokenString
}

//生成token
func GenerateToken(openid string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"openid": openid,                                                   //openid
		"iss":    "iris_bysj",                                              //签发者
		"iat":    time.Now().Unix(),                                        //签发时间
		"jti":    "9527",                                                   //jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
		"exp":    time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(), //过期时间
	})

	tokenString, _ := token.SignedString([]byte(jwtKey))
	fmt.Println("签发时间：", time.Now().Unix())
	fmt.Println("到期时间：", time.Now().Add(10 * time.Hour * time.Duration(1)).Unix())
	return tokenString
}
