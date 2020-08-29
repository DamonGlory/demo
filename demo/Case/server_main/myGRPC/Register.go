package myGRPC

import (
	"Damon.com/Case/server_main/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "Damon.com/Case/server_main/protobuf"
)

func RegisterGRPC(ctx *gin.Context,cfg *utils.Conf)(int,string){
	fmt.Println(cfg.MysqlgRPC.Address)
	conn, err := grpc.Dial(cfg.MysqlgRPC.Address,grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()
	//初始化客户端
	c:=pb.NewRegisterServiceClient(conn)
	//c:=pb.NewRegisterServiceRPCClient(conn)
	//调用方法
	req:=&pb.RegisterRequest{}
	req.Name=ctx.PostForm("username")
	req.Password = ctx.PostForm("password")
	fmt.Println("====================",req.Name)
	resp,err:=c.Register(context.Background(),req)
	if err != nil {
		return utils.RECODE_NODATA,utils.GetRecodeText(utils.RECODE_NODATA)
	}
	return int(resp.Recode),resp.Message
}