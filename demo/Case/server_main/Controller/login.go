package Controller

import (
	"Damon.com/Case/server_main/interceptor"
	"Damon.com/Case/server_main/log"
	"Damon.com/Case/server_main/myGRPC"
	"Damon.com/Case/server_main/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)
type  LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func Login(cfg *utils.Conf) gin.HandlerFunc {
	return func(c *gin.Context) {
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		//body,_ := ioutil.ReadAll(c.Request.Body)
		//logindata := LoginData{}
		//json.Unmarshal(body,&logindata)
		//username:=logindata.Username
		//password:=logindata.Password
		fmt.Printf("Login======%#v",username)
		flag:=interceptor.IsLogin(c,cfg)
		if flag{
			//重复登陆
			//TODO Send To kafka
			log.SetLog("damon",username,"重复登陆")
			c.JSON(200, gin.H{"msg":"重复登陆","address":utils.Cfg.ChatServer.Address})
		}else{
			//执行登录操作
			fmt.Println("===#####",username,password,"====3###")
			recode,msg:=myGRPC.LoginGRPC(username,password,cfg)
			if recode==0{
				//用户名密码正确
				//生成token
				tokenstr:=interceptor.SetToken(c,cfg)
				myGRPC.SetToken(username,tokenstr,cfg)
				//TODO 返回token
				log.SetLog("damon",username,"登录成功")
				c.JSON(200, gin.H{"msg":msg,"token":tokenstr,"address":utils.Cfg.ChatServer.Address})
				//登陆成功，send to kafka
			}else{
				c.JSON(200, gin.H{"msg":"用户名密码错误"})
			}
		}
	}
}

