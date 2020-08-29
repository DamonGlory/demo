package myGRPC

import (
	pb "Damon.com/Case/server_main/protobuf"
	"Damon.com/Case/server_main/utils"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func LoginGRPC(username,password string,cfg *utils.Conf)(int,string) {
	conn, err := grpc.Dial(cfg.MysqlgRPC.Address,grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()
	//初始化客户端
	c:=pb.NewLoginServiceClient(conn)
	//c:=pb.NewLoginServiceRPCClient(conn)
	//调用方法
	req:=&pb.LoginRequest{}
	req.Name=username
	req.Password = password
	fmt.Println("LoginGRPC:",req)
	resp,err:=c.Login(context.Background(),req)
	if err != nil {
		return utils.RECODE_NODATA,utils.GetRecodeText(utils.RECODE_NODATA)
	}
	return int(resp.Recode),resp.Message
}