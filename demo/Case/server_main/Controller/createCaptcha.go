package Controller

import (
	"Damon.com/Case/server_main/interceptor"
	"bytes"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateCaptcha() gin.HandlerFunc{
	return func(c *gin.Context) {
		//设置图片的默认宽高
		w, h := 400, 300
		captcha := captcha.NewLen(6)
		session := sessions.Default(c)
		session.Set(interceptor.RegisterCaptCha, captcha)
		session.Save()
		Serve(c.Writer, c.Request, captcha, ".png", "zh", false, w, h)
	}
}
func Serve(resp http.ResponseWriter, req *http.Request, id, ext, language string, download bool, w, h int) error {
	//设置验证码的格式
	resp.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	resp.Header().Set("Pragma", "no-cache")
	resp.Header().Set("Expires", "0")
	//反序列化
	fmt.Println(id)
	fmt.Printf("%#v\n",captcha.NewMemoryStore(captcha.CollectNum,captcha.Expiration).Get(id,false))
	//定义字节流
	var content bytes.Buffer
	switch ext {
	case ".png":
		//设置response的header格式
		resp.Header().Set("Content-Type", "image/png")
		//将图片放入content字节流
		//WriteImage(w io.Writer, id string, width, height int) error
		captcha.WriteImage(&content, id, w, h)
	case ".wav":
		resp.Header().Set("Content-Type", "audio/x-wav")
		resp.Header().Set("charset", "utf-8")
		//将音频放入content字节流
		// WriteAudio(w io.Writer, id string, lang string) error
		captcha.WriteAudio(&content, id, language)
	default:
		return captcha.ErrNotFound
	}
	//是否支持下载
	if download {
		resp.Header().Set("Content-Type", "application/octet-stream")
	}
	//将content发送给客户端
	//ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)
	http.ServeContent(resp, req, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}