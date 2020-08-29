package myGRPC

import (
	"Damon.com/Case/server_main/utils"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "Damon.com/Case/server_main/protobuf"
)

func SetToken(username,tokenStr string,cfg *utils.Conf) bool{
	fmt.Println(cfg.MysqlgRPC.Address)
	conn, err := grpc.Dial(cfg.MysqlgRPC.Address,grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()
	//初始化客户端
	c:=pb.NewSetTokenServiceClient(conn)
	//c:=pb.NewSetTokenServiceRPCClient(conn)
	req:=&pb.SetTokenRequest{}
	req.Name=username
	req.Token=tokenStr
	resp,err:=c.SetToken(context.Background(),req)
	fmt.Println("---------",resp)
	if err != nil {
		return false
	}
	if int(resp.Recode)==0{
		return true
	}else{
		return false
	}
}

//获取redis中的token
func GetToken(username string,cfg *utils.Conf) string{
	fmt.Println(cfg.MysqlgRPC.Address)
	conn, err := grpc.Dial(cfg.MysqlgRPC.Address,grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()
	//初始化客户端
	c:=pb.NewGetTokenServiceClient(conn)
	//c:=pb.NewGetTokenServiceRPCClient(conn)
	req:=&pb.GetTokenRequest{}
	req.Name=username
	resp,err:=c.GetToken(context.Background(),req)
	fmt.Println("---------",resp)
	if err != nil {
		return ""
	}
	if int(resp.Recode)==0{
		return resp.Token
	}else{
		return ""
	}
	return ""
}
