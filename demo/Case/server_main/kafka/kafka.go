package kafka

import (
	"Damon.com/Case/server_main/etcd"
	"Damon.com/Case/server_main/utils"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)
type Mgo struct {
	Database string `json:"database"`
	Collection string `json:"collection"`
}
type MgoData struct {
	//UserId int `json:"userid"`
	UserName string `json:"username"`
	Msg string `json:"msg"`
	Time time.Time `json:"time"`
}
//TODO 构建kafkaMSG
type LogMgoData struct {
	Topic string
	KafkaMgomsg
}
type KafkaMgomsg struct {
	Mgo
	MgoData
}
//
//
var KafkaProduce sarama.SyncProducer
var KafkaErr error
func init(){
	//连接kafka
	config:=sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll
	config.Producer.Partitioner=sarama.NewRandomPartitioner
	config.Producer.Return.Successes=true
	KafkaProduce,KafkaErr=sarama.NewSyncProducer([]string{utils.Cfg.Kafka.Address},config)
	if KafkaErr!=nil{
		fmt.Println("NewSyncProducer failed,KafkaErr:",KafkaErr)
		return
	}
	return
}
//
//
////构建LogDataChan
func SendToChan(mgo *etcd.EtcdMgo,username string,data string,time time.Time)(chan *LogMgoData){
	msg:=&LogMgoData{
		Topic:mgo.Topic,
	}
	msg.Database=mgo.Database
	msg.Collection=mgo.Collection
	msg.MgoData.UserName=username
	msg.MgoData.Msg=data
	msg.MgoData.Time=time
	//初始化LogDataChan
	LogDataChan := make(chan *LogMgoData,utils.Cfg.Kafka.Chan_max_size)
	LogDataChan <-msg
	return LogDataChan
}
//
////发送消息给kafka
func SendMessage(LogDataChan chan *LogMgoData)error{
		select {
		case ld := <-LogDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.Topic
			msgData := KafkaMgomsg{}
			msgData.Mgo=ld.Mgo
			msgData.MgoData=ld.MgoData
			//json编码kafkamsg
			data,err:=json.Marshal(msgData)
			if err!=nil{
				KafkaErr = err
				return KafkaErr
			}
			msg.Value = sarama.StringEncoder(string(data))
//			//fmt.Printf("%#v\n",msg.Value)
//			//发送消息到kafka
			go func(){
				pid, offset, err := KafkaProduce.SendMessage(msg)
				fmt.Println("1111111")
				if err != nil {
					fmt.Println("send message to kafka fail,err:", err)
					KafkaErr= err
					return
				}
				fmt.Printf("pid:%v,offset:%v", pid, offset)
			}()
		default:
			time.Sleep(time.Second)
		}
	return KafkaErr
}
