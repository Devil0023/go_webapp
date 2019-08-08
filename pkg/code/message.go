package code

func GetMsg(code int) string {

	msg, ok := MsgMap[code]

	if ok {
		return msg
	}

	return MsgMap[ERROR]
}
