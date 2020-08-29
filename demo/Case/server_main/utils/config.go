package utils

import "gopkg.in/ini.v1"

type Base struct {
	Host string `ini:"host"`
}
type MysqlgRPC struct {
	Address string `ini:"address"`
}
type Kafka struct {
	Address string `ini:"address"`
	Chan_max_size int `ini:"chan_max_size"`
}
type Etcd struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key string `ini:"key"`
}
type ChatServer struct {
	Address string `ini:"address"`
}

type Conf struct {
	Base
	MysqlgRPC
	Kafka
	Etcd
	ChatServer
}

var Cfg *Conf
func init(){
	Cfg = new(Conf)
	err:=ini.MapTo(Cfg,"./Conf/conf.ini")
	if err!=nil {
		panic(err)
	}
}
func GetConfig() *Conf{
	cfg := new(Conf)
	err:=ini.MapTo(cfg,"./Conf/conf.ini")
	if err!=nil {
		panic(err)
	}
	return cfg
}