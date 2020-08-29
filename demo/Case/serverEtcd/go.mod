module Damon.com/Case/serverEtcd

go 1.14

require (
	github.com/Shopify/sarama v1.27.0
	github.com/coreos/etcd v3.3.22+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	go.etcd.io/etcd v3.3.22+incompatible
	go.mongodb.org/mongo-driver v1.4.0
	go.uber.org/zap v1.15.0 // indirect
	google.golang.org/grpc v1.31.0 // indirect
	gopkg.in/ini.v1 v1.58.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
