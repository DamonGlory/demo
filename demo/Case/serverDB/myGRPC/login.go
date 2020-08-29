package myGRPC

import (
	"Damon.com/Case/serverDB/dao"
	"Damon.com/Case/serverDB/model"
	"Damon.com/Case/serverDB/protobuf"
	"Damon.com/Case/serverDB/utils"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
)
type LoginService struct {
	DB *gorm.DB
}
func (r LoginService) Login(cxt context.Context, res *protobuf.LoginRequest) (*protobuf.LoginResponse, error) {
	//执行Register方法
	u := model.User{}
	//数据库查询
	u.Name = res.Name
	u.Password = res.Password
	fmt.Println("Login res", res)
	recode := dao.Login(r.DB, &u)
	fmt.Println("myGRPC Login")
	resp := new(protobuf.LoginResponse)
	resp.Recode = int32(recode)
	resp.Message = utils.GetRecodeText(recode)
	return resp, nil
}

