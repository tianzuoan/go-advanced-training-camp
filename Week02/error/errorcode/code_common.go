package errorcode

const (
	SUCCESS        = 0
	FAIL           = 885800
	INVALID_PARAMS = 885801
	UNAUTH         = 885802
	NOT_FOUND      = 885803
	DB_ERR         = 885804
	UNKNOWN        = 885805
)

func init() {
	dict := map[int]string{
		SUCCESS:        "成功",
		FAIL:           "内部错误",
		INVALID_PARAMS: "非法请求参数",
		UNAUTH:         "无访问权限",
		NOT_FOUND:      "找不到资源",
		DB_ERR:         "数据库出错",
		UNKNOWN:        "未知错误",
	}
	RegisterErrMsg(dict)
}
