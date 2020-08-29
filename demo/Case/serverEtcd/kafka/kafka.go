package kafka

import (
	"Damon.com/Case/serverEtcd/log"
	"Damon.com/Case/serverEtcd/utils"
	"fmt"
	"github.com/Shopify/sarama"
)

var KafkaConsumer sarama.Consumer
var KafkaErr error

func init() {
	KafkaConsumer, KafkaErr = sarama.NewConsumer([]string{utils.Cfg.Kafka.Address}, nil)
	return
}

func GetMsgByTopic(topic string) ([]byte, error) {
	partitionList, err := KafkaConsumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return []byte(""), err
	}
	ch := make(chan int,len(partitionList))
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者

		pc, err := KafkaConsumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			return []byte(""), err
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(spc sarama.PartitionConsumer,ch1 chan int) {
			for msg := range spc.Messages() {
				//TODO
				//存盘
				//value, _ := utils.GetKafkaData(string(msg.Value))
				//log.LogToFile(value)
				value, _ := utils.GetKafkaMgoData(string(msg.Value))
				log.LogToMongoDB(value)
				fmt.Println(value.Time.Format("2006-01-02 15:04:05"))
			}
			ch1 <- 1
		}(pc,ch)
		//ch <- 1
	}
	<-ch
	return nil, nil
}
