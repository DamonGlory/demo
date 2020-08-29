package etcd

import (
	"Damon.com/Case/serverEtcd/models"
	"Damon.com/Case/serverEtcd/utils"
	"context"
	"fmt"
	"time"
	"go.etcd.io/etcd/clientv3"
)

//获取etcd的value
var EtcdCli *clientv3.Client
var EtcdErr error
var GetLogEntryCof []*models.LogEntry
var Key string
//func init(){
//	address:=utils.Cfg.Etcd.Address
//	timeout:=time.Duration(utils.Cfg.Etcd.Timeout)*time.Second
//	EtcdCli,EtcdErr =clientv3.New(clientv3.Config{
//		Endpoints:[]string{address},
//		DialTimeout:timeout,
//	})
//	//根据key获取value
//	Key = utils.Cfg.Etcd.Key
//	GetLogEntryCof = GetConf(Key)
//}
func GetConf(key string) (logEntryCof *models.LogEntry ){
	address:=utils.Cfg.Etcd.Address
	timeout:=time.Duration(utils.Cfg.Etcd.Timeout)*time.Second
	EtcdCli,EtcdErr =clientv3.New(clientv3.Config{
		Endpoints:[]string{address},
		DialTimeout:timeout,
	})
	ctx,cancel:=context.WithTimeout(context.Background(),6*time.Second)
	resp,err:=EtcdCli.Get(ctx,key)
	cancel()
	if err!=nil{
		EtcdErr = err
		return nil
	}
	for _,ev:=range resp.Kvs {
		logEntryCof,err=utils.GetEtcdValue(ev.Value)
		fmt.Println(logEntryCof)
		if err!=nil{
			EtcdErr = err
			return nil
		}
	}
	return
}
