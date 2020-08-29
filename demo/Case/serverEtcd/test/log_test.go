package test

import (
	"Damon.com/Case/serverEtcd/dao"
	"Damon.com/Case/serverEtcd/log"
	"testing"
	"time"
)

func TestLogToMongoDB(t *testing.T){
	logData := &dao.LogData{}
	logData.Mgo.Database="damon"
	logData.Mgo.Collection="log"
	logData.MgoData.UserId="123"
	logData.MgoData.UserName="damon"
	logData.MgoData.Time = time.Now()
	logData.MgoData.Msg="床前明月光，疑似地上霜"
	log.LogToMongoDB(logData)
}