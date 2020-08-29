package main

import (
	"Damon.com/Case/chatClis/chat"
	"Damon.com/chatCli/login"
	"fmt"
	"os"
)
var username string
var password string
var token string
var ch =make(chan int)
//func init(){
//	username =os.Args[1]
//	password =os.Args[2]
//	login.Login(username,password)
//}
func main(){
	//fmt.Printf("%#v\n",login.ServerData)
	username =os.Args[1]
	password =os.Args[2]
	login.Login(username,password)
	token=login.ServerData.Token
	if login.Login(username,password,token){
		fmt.Println("已经登录过了")
		//开始聊天
		go chat.Chat(username)
	}else{
		if login.Login(username,password){
			//开始聊天
			go chat.Chat(username)
		}else{
			fmt.Println("登录失败")
		}
	}
	<-ch
}
