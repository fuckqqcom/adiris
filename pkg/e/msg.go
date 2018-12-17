package e

var msgDict = map[int]string{
	Success: "响应正常",
	Error:   "服务器异常",
}

func GetMsg(code int) string {
	msg, ok := msgDict[code]
	if ok {
		return msg
	}

	return msgDict[Error]
}
