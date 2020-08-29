这是我写的一个简单的demo

消息中间件：kafka

数据库：存放个人信息（mysql）,存放session/token(redis),存放日志/评论（mongoDB）

分布式事务：etcd,grpc

涉及三个简单业务逻辑service(sys,DBservice,logservice)