package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

//type mgo struct {
//	database   string
//	collection string
//}
//func NewMgo(database, collection string) *mgo {
//	return &mgo{
//		database,
//		collection,
//	}
//}
func CheckErr(err error,str string){
	if err!=nil{
		panic(fmt.Errorf(str,err))
	}
}

//固化日志到mongoDB
func LogToMongo(logdata *LogData){
	//获取连接池
	mgo:=NewMgo("damon","log")
	mgo.InsertOne(logdata)
}

// 查询单个
func (m *Mgo) FindOne(key string, value interface{}) *mongo.SingleResult {
	client := DB.Mongo
	collection, _ := client.Database(m.Database).Collection(m.Collection).Clone()
	//collection.
	filter := bson.D{{key, value}}
	singleResult := collection.FindOne(context.TODO(), filter)
	return singleResult
}

//插入单个
func (m *Mgo) InsertOne(value interface{}) *mongo.InsertOneResult {
	client := DB.Mongo
	collection := client.Database(m.Database).Collection(m.Collection)
	insertResult, err := collection.InsertOne(context.TODO(), value)
	if err != nil {
		fmt.Println(err)
	}
	return insertResult
}

//查询集合里有多少数据
func (m *Mgo) CollectionCount() (string, int64) {
	client := DB.Mongo
	collection := client.Database(m.Database).Collection(m.Collection)
	name := collection.Name()
	size, _ := collection.EstimatedDocumentCount(context.TODO())
	return name, size
}

//按选项查询集合 Skip 跳过 Limit 读取数量 sort 1 ，-1 . 1 为最初时间读取 ， -1 为最新时间读取
func (m *Mgo) CollectionDocuments(Skip, Limit int64, sort int) *mongo.Cursor {
	client := DB.Mongo
	collection := client.Database(m.Database).Collection(m.Collection)
	SORT := bson.D{{"_id", sort}} //filter := bson.D{{key,value}}
	filter := bson.D{{}}
	findOptions := options.Find().SetSort(SORT).SetLimit(Limit).SetSkip(Skip)
	//findOptions.SetLimit(i)
	temp, _ := collection.Find(context.Background(), filter, findOptions)
	return temp
}

//获取集合创建时间和编号
func (m *Mgo) ParsingId(result string) (time.Time, uint64) {
	temp1 := result[:8]
	timestamp, _ := strconv.ParseInt(temp1, 16, 64)
	dateTime := time.Unix(timestamp, 0) //这是截获情报时间 时间格式 2019-04-24 09:23:39 +0800 CST
	temp2 := result[18:]
	count, _ := strconv.ParseUint(temp2, 16, 64) //截获情报的编号
	return dateTime, count
}

//删除文章和查询文章
func (m *Mgo) DeleteAndFind(key string, value interface{}) (int64, *mongo.SingleResult) {
	client := DB.Mongo
	collection := client.Database(m.Database).Collection(m.Collection)
	filter := bson.D{{key, value}}
	singleResult := collection.FindOne(context.TODO(), filter)
	DeleteResult, err := collection.DeleteOne(context.TODO(), filter, nil)
	if err != nil {
		fmt.Println("删除时出现错误，你删不掉的~")
	}
	return DeleteResult.DeletedCount, singleResult
}

//删除文章
func (m *Mgo) Delete(key string, value interface{}) int64 {
	client := DB.Mongo
	collection := client.Database(m.Database).Collection(m.Collection)
	filter := bson.D{{key, value}}
	count, err := collection.DeleteOne(context.TODO(), filter, nil)
	if err != nil {
		fmt.Println(err)
	}
	return count.DeletedCount

}

//删除多个
func (m *Mgo) DeleteMany(key string, value interface{}) int64 {
	client := DB.Mongo
	collection := client.Database(m.Database).Collection(m.Collection)
	filter := bson.D{{key, value}}

	count, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}
	return count.DeletedCount
}