package dao

import (
	"Damon.com/Case/serverDB/model"
	"Damon.com/Case/serverDB/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

func Login(db *gorm.DB, user *model.User) (int){
	var u model.User
	//注册前查询是否合格
	//判断用户名是否存在
	db.Where("name=? and password=?",user.Name,user.Password).First(&u)
	//不合格，返回错误代码recode
	fmt.Println("[",u,"]")
	if u.ID==0{
		fmt.Println("用户名密码错误")
		return utils.RECODE_LOGINERR
	}else{
		fmt.Println("登录成功")
		return utils.RECODE_OK
	}
}