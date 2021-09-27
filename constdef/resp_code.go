package constdef

const (
	//正常
	RespCodeOk int = 0
	//请求参数错误
	RespCodeRequestParamErr int = 1
	//数据库错误
	RespCodeDBErr int = 2
	//用户未登录
	RespCodeNoLoginErr int = 3
	//用户名重复
	RespCodeUserRepeatErr int = 4
	//用户不存在
	RespCodeUserNoExist int = 5
	//服务器内部错误
	RespCodeServerErr int = 6
)

func BuildServerErr() (int, string) {
	return RespCodeServerErr, "服务器内部错误"
}

func BuildRequestParamErr() (int, string) {
	return RespCodeRequestParamErr, "请求参数错误"
}

func BuildDBErr() (int, string) {
	return RespCodeDBErr, "数据库错误"
}

func BuildNoLoginErr() (int, string) {
	return RespCodeNoLoginErr, "用户未登录"
}

func BuildUserRepeatErr() (int, string) {
	return RespCodeUserRepeatErr, "用户名重复"
}
func BuildUserNotExist() (int, string) {
	return RespCodeUserNoExist, "账户或密码错误"
}
