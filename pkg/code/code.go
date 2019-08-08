package code

const (
	SUCCESS        = 0
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)

var MsgMap = map[int]string{

	SUCCESS:        "success",
	ERROR:          "error",
	INVALID_PARAMS: "invalid_params",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "ERROR_AUTH_CHECK_TOKEN_FAIL",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "ERROR_AUTH_CHECK_TOKEN_TIMEOUT",
	ERROR_AUTH_TOKEN:               "ERROR_AUTH_TOKEN",
	ERROR_AUTH:                     "ERROR_AUTH",
}
