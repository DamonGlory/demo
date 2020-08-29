package utils

import "gopkg.in/ini.v1"

type Mysql struct {
	Address string `ini:"address"`
	Database string `ini:"database"`
	Username string `ini:"username"`
	Password string `ini:"password"`

}
type Base struct {
	Address string `ini:"address"`
}
type Redis struct {
	Address string `ini:"address"`
}
type Conf struct {
	Mysql
	Base
	Redis
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