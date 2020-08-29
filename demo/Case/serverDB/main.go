package main

import (
	"Damon.com/Case/serverDB/model"
	"Damon.com/Case/serverDB/myGRPC"
	"Damon.com/Case/serverDB/protobuf"
	"Damon.com/Case/serverDB/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

var cfg *utils.Conf
var DB *gorm.DB
var err error
func init(){
	cfg = utils.GetConfig()
	openMysqlStr:=getMysql(cfg)
	DB,err=gorm.Open("mysql",openMysqlStr)
	if err!=nil{
		panic(err)
	}
	model.CreateTable(DB)
}

func main(){
	fmt.Println("main")
	//开启监听
	listen, err := net.Listen("tcp",cfg.Base.Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	//实例化server
	s:=grpc.NewServer()
	//在server上注册服务
	protobuf.RegisterRegisterServiceServer(s,&myGRPC.RegisterService{DB:DB})
	protobuf.RegisterLoginServiceServer(s,&myGRPC.LoginService{DB:DB})
	protobuf.RegisterSetTokenServiceServer(s,&myGRPC.SetTokenService{})
	protobuf.RegisterGetTokenServiceServer(s,&myGRPC.GetTokenService{})
	s.Serve(listen)
}

func getMysql(cfg *utils.Conf) string{
	address:=cfg.Mysql.Address
	database :=cfg.Mysql.Database
	username:=cfg.Mysql.Username
	password:=cfg.Mysql.Password
	str:=fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",username,password,address,database)
	return str
}