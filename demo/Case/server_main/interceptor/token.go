package interceptor

import (
	"Damon.com/Case/server_main/myGRPC"
	"Damon.com/Case/server_main/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var mySigningKey = []byte("www.Damon.com")

type Claims struct {
	UserID string
	jwt.StandardClaims
}
//token操作
func IsLogin(c *gin.Context,cfg *utils.Conf) bool{
	tokenStr:=c.GetHeader("Authorization")
	username:=c.PostForm("username")
	//首次登录
	if tokenStr == "" {
		tokenStr =SetToken(c,cfg)
		if 	myGRPC.SetToken(username,tokenStr,cfg){
			fmt.Println(username)
		}
		return false
	}
	str:=myGRPC.GetToken(username,cfg)
	//fmt.Println("====tokenStr=====",tokenStr)
	//fmt.Println("====str=====",str)
	//fmt.Println(strings.TrimPrefix(tokenStr,"Bearer ")==str)
	if strings.TrimPrefix(tokenStr,"Bearer ")==str{
	//if tokenStr==str{
		//相同  返回true
		fmt.Println("true======")
		return true
	}else{
		//不同  返回false
		fmt.Println("false======")
		return false
	}
}


//token是否有效
func ParseToken(c *gin.Context,tokenStr string) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(c.Request,request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	return token,err
}
//生成token
func SetToken(c *gin.Context,cfg *utils.Conf ) string{
	//获取用户名密码
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	//1.Create the Claims
	/**
	type StandardClaims struct {
		Audience  string `json:"aud,omitempty"` 用户（受众）
		ExpiresAt int64  `json:"exp,omitempty"` 到期时间(Unix()格式）
		Id        string `json:"jti,omitempty"` 唯一标识（编号）
		IssuedAt  int64  `json:"iat,omitempty"` 签发时间(Unix()格式）
		Issuer    string `json:"iss,omitempty"` 颁发者
		NotBefore int64  `json:"nbf,omitempty"` 生效时间(因为是int64格式，所以用Unix()格式）
		Subject   string `json:"sub,omitempty"` 主题
	}
	*/
	userId:=username+password
	standardClaims:=jwt.StandardClaims{
		ExpiresAt:time.Now().Add(time.Minute*60).Unix(), //在一个小时后失效
		IssuedAt:time.Now().Unix(),
		Issuer:cfg.Base.Host,
		Subject:"userToken",
	}
	claims:=&Claims{
		UserID:userId,
		StandardClaims:standardClaims,
	}
	//2.加密生成token结构体
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	//密钥签名
	tokenStr,err:=token.SignedString(mySigningKey)
	if err!=nil{
		fmt.Println("token SignedString failed,err:",err)
		return ""
	}
	return tokenStr
}