package dao

import (
	"Damon.com/Case/serverDB/model"
	"Damon.com/Case/serverDB/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

//注册
func Register(db *gorm.DB, user *model.User) (int){
	var u model.User
	//注册前查询是否合格
	//判断用户名是否存在
	db.Where("name=?",user.Name).First(&u)
	//不合格，返回错误代码recode
	if u.ID!=0{
		fmt.Println("数据库中存在，无效用户名，不可以注册")
		return utils.RECODE_REGISTEREXIST
	}else{
		//合格,数据库添加，返回recode,msg
		fmt.Println(user.Name)
		fmt.Println(user.Name,user.Password)
		fmt.Println("数据库中不存在，有效用户名，可以注册")
		db.Create(user)
		return utils.RECODE_REGISTEROK
	}
}