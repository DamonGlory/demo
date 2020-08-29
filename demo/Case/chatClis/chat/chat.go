package chat

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

var ch =make(chan int)
func CHandleErr(err error,why string){
	if err!=nil{
		fmt.Println(why,err)
		os.Exit(1)
	}
}
func Chat(username string){
	//在命令行参数携带昵称
	fmt.Println(username)
	conn,err:=net.Dial("tcp","127.0.0.1:8898")
	CHandleErr(err,"dail err")
	defer conn.Close()
	go handleSend(conn,username)
	go handleReceive(conn)
	<-ch
}
func handleReceive(conn net.Conn) {
	buffer:=make([]byte,1024)
	for{

		n,err:=conn.Read(buffer)
		if err!=io.EOF{
			CHandleErr(err,"conn read")
		}
		if n>0{
			msg:=string(buffer[:n])
			fmt.Println(msg)
		}
	}
	
}
func handleSend(conn net.Conn,name string) {
	_, err := conn.Write([]byte(name))
	CHandleErr(err,"conn write")
	//发送昵称到服务端
	reader := bufio.NewReader(os.Stdin)
	for{
		//读取标准输入
		lineBytes, _, _ := reader.ReadLine()
		//发送到服务端
		_,err:=conn.Write(lineBytes)
		if string(lineBytes)=="exit"{
			break
		}
		CHandleErr(err,"conn write")
	}
}

