package main

import (
	"Damon.com/Case/serverEtcd/etcd"
	"Damon.com/Case/serverEtcd/kafka"
	"Damon.com/Case/serverEtcd/utils"
)

func main(){
	//获取etcd的value
	logEntryCof:=etcd.GetConf(utils.Cfg.Etcd.Key)
	go kafka.GetMsgByTopic(logEntryCof.Topic)
	select {

	}
}
