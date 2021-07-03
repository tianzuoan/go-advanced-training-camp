package errorcode

import (
	"log"
)

var errMsgDict = make(map[int]string)

func RegisterErrMsg(dict map[int]string) {
	for code, errMsg := range dict {
		if errMsgDict[code] != "" {
			log.Fatalf("错误码初始化错误,重复定义的code:%d", code)
		}
		errMsgDict[code] = errMsg
	}
}

func GetErrMsg(code int) string {
	return errMsgDict[code]
}
