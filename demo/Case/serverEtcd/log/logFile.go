package log

import (
	"Damon.com/Case/serverEtcd/dao"
)
//日志写入日志文件
//func LogToFile(kafkaMsg *models.Kafkamsg){
//	filename:=kafkaMsg.Filename
//	data:=kafkaMsg.Data
//	username:=kafkaMsg.Username
//	t:= kafkaMsg.Time
//	file, err :=os.OpenFile("./logFile/"+filename,os.O_CREATE|os.O_RDWR|os.O_APPEND,0666)
//	fmt.Println("./log/"+filename)
//	if err != nil {
//		fmt.Println("open file failed, err:", err)
//		return
//	}
//	defer file.Close()
//	str:=fmt.Sprintf("[%s]:[%s] %s\n",t.Format("2006-01-02 15:04:05"),username,data)
//	file.WriteString(str)
//}
//日志写入mongoDB
func LogToMongoDB(logdata *dao.LogData){
	mgo:=dao.NewMgo(logdata.Mgo.Database,logdata.Mgo.Collection)
	mgo.InsertOne(logdata.MgoData)
}