package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

type Client struct {
	conn net.Conn
	name string
	addr string
}
var (
	//用昵称做键
	clientMap = make(map[string]Client)
)
func SHandleErr(err error,why string){
	if err!=nil{
		fmt.Println(err,why)
		os.Exit(1)
	}
}
func main(){
	listen,err:=net.Listen("tcp","127.0.0.1:8898")
	SHandleErr(err,"net listen")
	defer func() {
		for _,c:=range clientMap{
			c.conn.Write([]byte("服务器维护"))
		}
		listen.Close()
	}()
	for{
		conn,err:=listen.Accept()
		SHandleErr(err,"listen accept")
		//给已上线用户发送上线消息
		buffer:=make([]byte,1024)
		var clientName string
		for{
			n,err:=conn.Read(buffer)
			SHandleErr(err,"conn.Read(buffer)")
			if n>0{
				clientName=string(buffer[:n])
				break
			}

		}
		fmt.Println(clientName+"上线了")
		for _,c:=range clientMap{
			c.conn.Write([]byte(c.name+"上线了"))
		}
		client:=Client{conn:conn,name:clientName,addr:conn.RemoteAddr().String()}
		clientMap[clientName]=client
		go ioWithConn(client)
	}

}
func ioWithConn(client Client) {
	//clientAddr:=client.conn.RemoteAddr().String()
	buffer:=make([]byte,1024)
	for{
		n,err:=client.conn.Read(buffer)
		if err!=io.EOF{
			SHandleErr(err,"conn Read")
		}
		if n>0{
			msg:=string(buffer[:n])
			fmt.Printf("[%s] say :[%s]\n",client.name,msg)
			strs:=strings.Split(msg,"#")
			if len(strs)>1{
				fmt.Println(strs)
				targetName:=strs[0]
				targetMsg:=strs[1]
				if targetName =="all"{
					//群发消息
					for _,c:=range clientMap{
						c.conn.Write([]byte(c.name+":"+targetMsg))
					}
				}else{
					//点对点发消息
					for addr,c:=range clientMap{
						if addr==targetName{
							c.conn.Write([]byte(client.name+":"+targetMsg))
						}
					}
				}
			}else{
				if msg =="exit"{
					for name,c:=range clientMap{
						if c.conn==client.conn{
							delete(clientMap,name)
						}else{
							c.conn.Write([]byte(c.name+"下线了"))
						}
					}
				}else{
					client.conn.Write([]byte("已阅读："+msg))
				}
			}

		}
	}


}