package etcd
//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"go.etcd.io/etcd/clientv3"
//	"time"
//)
//
//
////Etcd中value的结构
//type LogEntry struct {
////	Mytype string `json:"mytype"`
////	Filename string `json:"filename"`
////	Topic string `json:"topic"`
////}
//var EtcdCli *clientv3.Client
//var GetLogEntryCof []*LogEntry
//var EtcdErr error
//var Key string
////func init(){
////	address:=utils.Cfg.Etcd.Address
////	timeout:=time.Duration(utils.Cfg.Etcd.Timeout)*time.Second
////	EtcdCli,EtcdErr =clientv3.New(clientv3.Config{
////		Endpoints:[]string{address},
////		DialTimeout:timeout,
////	})
////	//根据key获取value
////	Key = utils.Cfg.Etcd.Key
////	GetLogEntryCof = GetConf(Key)
////}
////etcd get(通过key获取value)
//func GetConf(key string) []*LogEntry{
//	var logEntryCof []*LogEntry
//	ctx,cancel:=context.WithTimeout(context.Background(),6*time.Second)
//	resp,err:=EtcdCli.Get(ctx,key)
//	cancel()
//	if err!=nil{
//		EtcdErr = err
//		return nil
//	}
//	for _,ev:=range resp.Kvs{
//		err=json.Unmarshal(ev.Value,&logEntryCof)
//		if err!=nil{
//			EtcdErr = err
//			return nil
//		}
//	}
//	return logEntryCof
//}
////etcd put
//func PutConf(){
////	value:=`[{"mytype":"register","filename":"./register.log","topic":"Damon_test0"},
////{"mytype":"register","filename":"./register.log","topic":"Damon_test1"},
////{"mytype":"register","filename":"./sql.log","topic":"Damon_test2"}]`
//	value:=`{"database":"damon","collection":"log","topic":"Damon_test0"}`
//	fmt.Println(value)
//}
//
//
