package Controller

import (
	"Damon.com/Case/server_main/interceptor"
	"Damon.com/Case/server_main/log"
	"Damon.com/Case/server_main/myGRPC"
	"Damon.com/Case/server_main/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

//注册
func Register(cfg *utils.Conf) gin.HandlerFunc {
	return func(c *gin.Context) {
		//验证验证码
		captchaFlag:=interceptor.CaptchaVerify(c)
		//验证码验证成功，正确处理
		if captchaFlag{
			registerResult(c, cfg)
		}else{
			registerErr(c)
		}
	}
}
func registerResult(c *gin.Context, cfg *utils.Conf) {
	//GRPC通讯数据库服务进程获取结果result
	recode, msg := myGRPC.RegisterGRPC(c, cfg)
	username:=c.PostForm("username")
	//TODO
	//send to kafka
	log.SetLog("damon",username,"注册成功")
	fmt.Println(username)
	c.JSON(utils.RECODE_OK, gin.H{"recode": recode, "message": msg})

}
func registerErr(c *gin.Context) {
	c.JSON(utils.RECODE_OK, gin.H{"message": "验证码错误"})
}
