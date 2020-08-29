package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"github.com/etcd-io/etcd/clientv3"
)

//watch监视etcd,然后更改配置
func Watch(etcdEndpoints []string,key string,newEtcdmgoch chan<- *EtcdMgo){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdEndpoints,
		DialTimeout: 2 * time.Second,
	})
	if nil != err {
		panic(err)
	}
	ch:=cli.Watch(context.Background(),key)
	//发现变化时，将最新数据写入chan
	for wresp :=range ch{
		for _,evt:=range wresp.Events{
			var newEtcdmgo *EtcdMgo
			err:=json.Unmarshal(evt.Kv.Value,&newEtcdmgo)
			if err!=nil{
				fmt.Println("json Unmarshal failed,err:",err)
				continue
			}
			newEtcdmgoch<-newEtcdmgo
		}
	}
}