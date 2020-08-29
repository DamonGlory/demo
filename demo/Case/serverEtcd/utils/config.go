package utils

import "gopkg.in/ini.v1"

type Kafka struct {
	Address string `ini:"address"`
	Chan_max_size int `ini:"chan_max_size"`
}
type Etcd struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key string `ini:"key"`
}
type MongoDB struct {
	Address string `ini:"address"`
}
type Conf struct {
	Kafka
	Etcd
	MongoDB
}


var Cfg *Conf
func init(){
	Cfg = new(Conf)
	err:=ini.MapTo(Cfg,"./Conf/conf.ini")
	if err!=nil {
		panic(err)
	}
}
