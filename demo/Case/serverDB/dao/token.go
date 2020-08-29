package dao

import (
	"Damon.com/Case/serverDB/utils"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func SetToken(username,token string) bool{
	RDB,err:=redis.Dial("tcp",utils.Cfg.Redis.Address)
	if err!=nil{
		panic(err)
	}
	defer RDB.Close()
	_, err = RDB.Do("Set", username, token)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}


func GetToken(username string)string{
	RDB,err:=redis.Dial("tcp",utils.Cfg.Redis.Address)
	if err!=nil{
		panic(err)
	}
	defer RDB.Close()
	r, err := redis.String(RDB.Do("Get", username))
	if err != nil {
		return ""
	}
	return r
}
