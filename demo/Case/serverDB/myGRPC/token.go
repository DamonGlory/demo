package myGRPC

import (
	"Damon.com/Case/serverDB/dao"
	"Damon.com/Case/serverDB/protobuf"
	"Damon.com/Case/serverDB/utils"
	"context"
)

type SetTokenService struct {

}
type GetTokenService struct {

}

func (r SetTokenService) SetToken(cxt context.Context,res *protobuf.SetTokenRequest)(*protobuf.SetTokenResponse, error){
	//TODO
	//dao.session
	flag:=dao.SetToken(res.Name,res.Token)
	resp := new(protobuf.SetTokenResponse)
	if flag{
		resp.Recode=utils.RECODE_OK
		resp.Message = utils.GetRecodeText(utils.RECODE_OK)
	}else{
		resp.Recode=utils.RECODE_SETTOKENERR
		resp.Message = utils.GetRecodeText(utils.RECODE_SETTOKENERR)
	}
	return resp,nil
}
func (r GetTokenService) GetToken(cxt context.Context,res *protobuf.GetTokenRequest)(*protobuf.GetTokenResponse, error){
	//TODO
	//dao.session
	str:=dao.GetToken(res.Name)
	resp := new(protobuf.GetTokenResponse)
	resp.Recode=utils.RECODE_OK
	resp.Message = utils.GetRecodeText(utils.RECODE_OK)
	resp.Token=str
	return resp,nil
}

