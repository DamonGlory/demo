package log

import (
	"Damon.com/Case/server_main/etcd"
	"Damon.com/Case/server_main/kafka"
	"fmt"
	"time"
)

//
//import (
//	"Damon.com/Case/server_main/etcd"
//	"Damon.com/Case/server_main/kafka"
//	"fmt"
//	"time"
//)
//
//func SetLog(mytype,username,data string,time time.Time)(err error){
//	//从etcd中获取value
//	loEntryCof:=etcd.GetLogEntryCof
//	if etcd.EtcdErr!=nil {
//		return etcd.EtcdErr
//	}
//	fmt.Printf("%#v",loEntryCof[0])
//	switch mytype {
//	case "login":
//		for _,ev:=range loEntryCof {
//			if ev.Mytype=="login"{
//				//send to kafka
//				fmt.Println(ev.Topic,ev.Filename,data)
//				logdataChan:=kafka.SendToChan(ev.Topic,ev.Filename,data,username,time)
//				fmt.Println("send to kafka login log1")
//				err=kafka.SendMessage(logdataChan)
//				fmt.Println("send to kafka login log2")
//			}
//		}
//	case "register":
//		for _,ev:=range loEntryCof {
//			if ev.Mytype=="register"{
//				//send to kafka
//				logdataChan:=kafka.SendToChan(ev.Topic,ev.Filename,data,username,time)
//				err=kafka.SendMessage(logdataChan)
//			}
//		}
//	default:
//		for _,ev:=range loEntryCof {
//			if ev.Mytype=="sql"{
//				//send to kafka
//				logdataChan:=kafka.SendToChan(ev.Topic,ev.Filename,data,username,time)
//				err=kafka.SendMessage(logdataChan)
//			}
//		}
//	}
//	return
//}
func SetLog(key string,username string,data string) (err error){
	//注册服务
	etcd.EtcdService()
	//获取etcd中的数据
	mgo := etcd.GetConf(key)
	if mgo==nil{
		fmt.Println("nil")
		return
	}
	fmt.Println("--------------",mgo)
	//TODO 准备向kafka中传输的数据
	logdataChan:=kafka.SendToChan(mgo,username,data,time.Now())
	err=kafka.SendMessage(logdataChan)
	return
}
