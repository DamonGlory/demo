package test

import (
	"Damon.com/Case/server_main/etcd"
	"testing"
)

func TestEtcdService(t *testing.T){
	etcd.EtcdService()
	newEtcdmgoch:=make(chan *etcd.EtcdMgo)
	etcd.Watch([]string{"127.0.0.1:2379"},"damon",newEtcdmgoch)

}