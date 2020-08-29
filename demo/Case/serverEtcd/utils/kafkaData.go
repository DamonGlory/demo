package utils

import (
	"Damon.com/Case/serverEtcd/dao"
	"Damon.com/Case/serverEtcd/models"
	"encoding/json"
)

func GetKafkaData(str string) (kafkaMsg *models.Kafkamsg,err error){
	err=json.Unmarshal([]byte(str),&kafkaMsg)
	return kafkaMsg,err
}
func GetKafkaMgoData(str string) (kafkaMgoData *dao.LogData,err error){
	err=json.Unmarshal([]byte(str),&kafkaMgoData)
	return kafkaMgoData,err
}