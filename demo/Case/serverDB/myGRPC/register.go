package myGRPC

import (
	"context"
	"fmt"

	"Damon.com/Case/serverDB/dao"
	"Damon.com/Case/serverDB/model"
	"Damon.com/Case/serverDB/protobuf"
	"Damon.com/Case/serverDB/utils"
	"github.com/jinzhu/gorm"
)

type RegisterService struct {
	DB *gorm.DB
}

func (r RegisterService) Register(cxt context.Context, res *protobuf.RegisterRequest) (*protobuf.RegisterResponse, error) {
	//执行Register方法
	u := model.User{}
	//数据库注册
	u.Name = res.Name
	u.Password = res.Password
	fmt.Println("Register res", res)
	recode := dao.Register(r.DB, &u)
	fmt.Println("myGRPC Register")
	resp := new(protobuf.RegisterResponse)
	resp.Recode = int32(recode)
	resp.Message = utils.GetRecodeText(recode)
	return resp, nil
}
