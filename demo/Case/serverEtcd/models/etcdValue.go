package models

//Etcd中value的结构
type LogEntry struct {
	Database string 	`json:"database"`
	Collection string 	`json:"collection"`
	Topic string 		`json:"topic"`
}
