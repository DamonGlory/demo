package interceptor

import (
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)
//session中的key
var RegisterCaptCha = "registerCaptCha"

//验证码
//1.生成一个k-v session
//2.生成验证码
//3.取参数，并于早先存于session中的验证码校对
//校验验证码
func CaptchaVerify(c *gin.Context) bool {
	//1.从session中取出验证码
	//1.1获取存放验证码的session
	session:=sessions.Default(c)
	//获取session中的验证码
	captchaID:=session.Get(RegisterCaptCha)
	if captchaID!=nil{
		//存在
		//验证一次后，需要更新验证码
		//a.删除原session
		session.Delete(RegisterCaptCha)
		//b.生成新session(抽象形容，验证码=session.value)
		session.Save()
		//2.将request中的code与session中的校对
		if captcha.VerifyString(captchaID.(string), c.PostForm("code")) {
			return true
		} else {
			return false
		}
		return true
	}else{
		//不存在
		return false
	}
}


//补充，根据验证码生成其他文件（图片等）
