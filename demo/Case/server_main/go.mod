module Damon.com/Case/server_main

go 1.14

require (
	github.com/Shopify/sarama v1.27.0
	github.com/coreos/etcd v3.3.22+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dchest/captcha v0.0.0-20170622155422-6a29415a8364
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/etcd-io/etcd v3.3.25+incompatible
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1 // indirect
	go.etcd.io/etcd v3.3.22+incompatible
	go.uber.org/zap v1.15.0 // indirect
	google.golang.org/grpc v1.31.0
	gopkg.in/ini.v1 v1.58.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
