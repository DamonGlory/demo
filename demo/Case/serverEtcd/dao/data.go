package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//kafka传递的日志数据(需要存入mongoDB的)
type Mgo struct {
	Database string `json:"database"`
	Collection string `json:"collection"`
}
func NewMgo(database, collection string) *Mgo {
	return &Mgo{
		database,
		collection,
	}
}
type MgoData struct {
	UserId string `json:"userid"`
	UserName string `json:"username"`
	Msg string `json:"msg"`
	Time time.Time `json:"time"`
}
type LogData struct {
	Mgo
	MgoData
}
type Database struct {
	Mongo  * mongo.Client
}
var DB *Database

//初始化
func init() {
	DB = &Database{
		Mongo: SetConnect(),
	}
}
// 连接设置
func SetConnect() *mongo.Client{
	//uri := "mongodb://"+Cfg.MongoDB.Address
	uri := "mongodb://127.0.0.1:27017"
	ctx ,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx,options.Client().ApplyURI(uri).SetMaxPoolSize(20)) // 连接池
	if err !=nil{
		fmt.Println(err)
	}
	return client
}