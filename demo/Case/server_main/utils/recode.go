package utils

const (
	RECODE_OK =0
	RECODE_NODATA = 4002
	RECODE_REGISTERERR = 4001
)

var recodeText = map[int]string{
	RECODE_OK : "成功",
	RECODE_NODATA:"连接数据库失败",
	RECODE_REGISTERERR : "注册失败",
}

func GetRecodeText(code int) string{
	return recodeText[code]
}