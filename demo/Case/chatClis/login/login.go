package login

import (
	"Damon.com/Case/chatClis/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type LoginData struct{
	Username string `json:"username"`
	Password string	`json:"password"`
}
var ServerData *utils.ServerRecode

//不带token登录  ---作废
/**
func LoginNoToke(username,password string){
	resp,err:=http.PostForm("http://127.0.0.1:8080/login",url.Values{"username":{username},"password":{password}})
	if err!=nil{
		fmt.Println(err)
	}
	body,err := ioutil.ReadAll(resp.Body)
	//serverRecode:= utils.ServerRecode{}
	json.Unmarshal(body,&ServerData)
	//fmt.Printf("%#v",ServerData)
}
*/
//再次登陆
/**
func IsLogin(username,password,token string)bool{
	client := &http.Client{}
	logindata:=LoginData{}
	logindata.Username=username
	logindata.Password=password
	buf,err:=json.Marshal(logindata)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%#v",string(buf))
	buffer:=bytes.NewReader(buf)
	req,err:=http.NewRequest("POST","http://127.0.0.1:8080/login",buffer)
	//resp,err:=http.PostForm(
	//	"http://127.0.0.1:8080/login",
	//	url.Values{"username":{username}, "password":{password}},
	//	)
	req.Header.Set("Authorization",token)
	resp,err:=client.Do(req)
	if err!=nil{
		fmt.Println(err)
	}
	var aServerData *utils.ServerRecode
	body,err := ioutil.ReadAll(resp.Body)
	//serverRecode:= utils.ServerRecode{}
	fmt.Println(token)
	json.Unmarshal(body,&aServerData)
	fmt.Println(aServerData)
	return true
}
*/
func Login(strs ...string)bool{
	client := &http.Client{}
	var token string
	logindata:=LoginData{}
	for k,v:=range strs{
		switch k {
		case 0:
			logindata.Username=v
		case 1:
			logindata.Password=v
		default:
			token = v
		}
	}
	fmt.Println(logindata.Username,logindata.Password,token)
	buf,err:=json.Marshal(logindata)
	if err!=nil{
		fmt.Println(err)
	}
	//fmt.Printf("%#v",string(buf))
	buffer:=bytes.NewReader(buf)
	req,err:=http.NewRequest("POST","http://127.0.0.1:8080/login",buffer)
	req.Header.Set("Authorization",token)
	resp,err:=client.Do(req)
	if err!=nil{
		fmt.Println(err)
	}
	body,err := ioutil.ReadAll(resp.Body)
	//serverRecode:= utils.ServerRecode{}
	json.Unmarshal(body,&ServerData)
	//fmt.Println(ServerData)
	if ServerData.Msg=="重复登陆"{
		return true
	}
	return false
}
