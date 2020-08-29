package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/etcd-io/etcd/clientv3"
	"sync"
	"time"
)
type EtcdMgo struct {
	Database string 	`json:"database"`
	Collection string 	`json:"collection"`
	Topic string 		`json:"topic"`
}
type Service struct {
	closeChan chan struct{}    //关闭通道
	client    *clientv3.Client //etcd v3 client
	leaseID   clientv3.LeaseID //etcd 租约id
	key       string           //键
	val       string           //值
	wg        sync.WaitGroup
}
// NewService 构造一个注册服务
func NewService(etcdEndpoints []string, key string, val string) (*Service, error) {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdEndpoints,
		DialTimeout: 2 * time.Second,
	})

	if nil != err {
		return nil, err
	}

	s := &Service{
		client:    cli,
		closeChan: make(chan struct{}),
		key:       key,
		val:       val,
	}

	return s, nil
}

// Start 开启注册
// @param - ttlSecond 租期(秒)
func (s *Service) Start(ttlSecond int64) error {
	// minimum lease TTL is 5-second
	resp, err := s.client.Grant(context.TODO(), ttlSecond)
	if err != nil {
		panic(err)
	}
	s.leaseID = resp.ID
	_, err = s.client.Put(context.TODO(), s.key, s.val, clientv3.WithLease(s.leaseID))
	if err != nil {
		panic(err)
	}
	ch, err1 := s.client.KeepAlive(context.TODO(), s.leaseID)
	if nil != err1 {
		panic(err)
	}
	fmt.Println("[discovery] Service Start leaseID:[%d] key:[%s], value:[%s]", s.leaseID, s.key, s.val)
	s.wg.Add(1)
	defer s.wg.Done()
	//确保随时可以关闭
	for{
		select {
		case <-s.closeChan: //如果chan关闭，revoke
			return s.revoke()
		case <-s.client.Ctx().Done():
			return errors.New("server closed")
		case ka, ok := <-ch:
			if !ok {
				fmt.Println("[discovery] Service Start keep alive channel closed")
				return s.revoke()
			} else {
				fmt.Println("[discovery] Service Start recv reply from Service: %s, ttl:%d", s.key, ka.TTL)
			}

		}
	}
	return nil
}
// Stop 停止
func (s *Service) Stop() {
	close(s.closeChan)
	s.wg.Wait()
	s.client.Close()
}
//租约释放
func (s *Service) revoke() error {

	_, err := s.client.Revoke(context.TODO(), s.leaseID)

	if err != nil {
		fmt.Println("[discovery] Service revoke key:[%s] error:[%s]", s.key, err.Error())
	} else {
		fmt.Println("[discovery] Service revoke successfully key:[%s]", s.key)
	}

	return err
}

//实现etcd的服务注册功能
func EtcdService(){
	//实例化client
	mgo := EtcdMgo{}
	mgo.Topic="damon"
	mgo.Database="damon"
	mgo.Collection="log"
	byte,_:=json.Marshal(mgo)
	s,err:=NewService([]string{"127.0.0.1:2379"},"damon",string(byte))
	if err!=nil{
		panic(err)
	}
	//s.client.Put(context.Background(),s.key,s.val)
	go s.Start(time.Second.Nanoseconds())
	//time.Sleep(time.Second*5)
	//s.Stop()
	fmt.Println("success")
}

//根据key获取etcd中的数据
func GetConf(key string) ( *EtcdMgo){
	cli,err :=clientv3.New(clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:5*time.Second,
	})
	ctx,cancel:=context.WithTimeout(context.Background(),6*time.Second)
	resp,err:=cli.Get(ctx,key)
	cancel()
	if err!=nil{
		fmt.Println("cli.Get",err)
		return nil
		}
	fmt.Println("resp===",resp)
	var etcdmgo *EtcdMgo
	for _,ev:=range resp.Kvs{
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)

		err=json.Unmarshal(ev.Value,&etcdmgo)
		if err!=nil{
			fmt.Println("json.Unmarshal",err)
			return nil
			}
		}
	fmt.Println("-2-2-2-2-2-",etcdmgo)
	return etcdmgo
}

