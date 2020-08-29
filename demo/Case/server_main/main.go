package main

import (
	"Damon.com/Case/server_main/interceptor"
	"Damon.com/Case/server_main/Controller"
	"Damon.com/Case/server_main/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)
var cfg *utils.Conf
func init(){
	//获取配置文件信息
	//cfg = utils.GetConfig()
	cfg=utils.Cfg
	//初始化ETCD
	//初始化kafka
}
func main(){
	//处理用户请求
	//fmt.Println(cfg.Base.Host)
	//注册
	r:=gin.Default()
	r.LoadHTMLGlob("./*.html")
	r.Use(interceptor.Session(interceptor.RegisterCaptCha))
	r.Handle("POST","/register",Controller.Register(cfg))
	r.Handle("POST","/login",Controller.Login(cfg))
	r.Handle("GET","/captcha",Controller.CreateCaptcha())
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run(cfg.Base.Host)
}
