package utils

const (
	RECODE_OK =0
	RECODE_REGISTEREXIST = 4001
	RECODE_REGISTEROK = 4000
	RECODE_LOGINERR = 4002
	RECODE_SETTOKENERR = 4003
	RECODE_GETTOKENERR = 4004
)

var recodeText = map[int]string{
	RECODE_OK : "成功",
	RECODE_REGISTEREXIST : "无效用户名",
	RECODE_REGISTEROK : "注册成功",
	RECODE_LOGINERR:"用户名密码错误",
	RECODE_SETTOKENERR:"token 设置失败",
	RECODE_GETTOKENERR:"token 获取失败",
}

func GetRecodeText(code int) string{
	return recodeText[code]
}