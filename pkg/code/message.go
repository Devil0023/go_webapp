package code

//GetMsg 获取错误码
func GetMsg(code int) string {

	msg, ok := MsgMap[code]

	if ok {
		return msg
	}

	return MsgMap[ERROR]
}
