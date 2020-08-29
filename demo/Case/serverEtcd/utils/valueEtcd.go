package utils

import (
	"Damon.com/Case/serverEtcd/models"
	"encoding/json"
)

func GetEtcdValue(buf []byte) (logEntryCof *models.LogEntry,err error){
	err=json.Unmarshal(buf,&logEntryCof)
	return
}