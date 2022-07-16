package errx

var errMessage map[uint32]string

func init() {
	errorMessage := make(map[uint32]string)

	errorMessage[SUCCESS] = "SUCCESS"
	errorMessage[SERVER_COMMON_ERROR] = "SERVER INTERNAL ERROR"
	errorMessage[REQ_PARAM_ERROR] = "REQUEST PARAMETER ERROR"
	errorMessage[TOKEN_EXPIRED_ERROR] = "TOKEN HAS BEEN EXPIRED"
	errorMessage[TOKEN_INVALID_ERROR] = "TOKEN HAS BEEN Invalid"
	errorMessage[TOKEN_GENERATE_ERROR] = "TOKEN GENERATE FAILED"
	errorMessage[DB_ERROR] = "DATABASE ERROR"
	errorMessage[DB_AFFECTED_ZERO_ERROR] = "DATABASE AFFECTED 0 rows"
	errorMessage[EMAIL_HAS_BEEN_REGISTERED] = "USER HAS BEEN REGISTERED"
	errorMessage[USER_NOT_EXIST] = "USER NOT EXIST"
	errorMessage[USER_PASSWORD_INCORRECT] = "USER PASSWORD INCORRECT"
	errorMessage[MOVIE_NOT_EXIST] = "MOVIE NOT EXIST"
	errorMessage[MOVIE_ALREADY_LIKED] = "MOVIE IS ALREADY LIKED"
	errorMessage[POST_NOT_EXIST] = "POST NOT EXIST"
	errorMessage[LIST_NOT_EXIST] = "LIST NOT EXIST"
}

func MapErrMsg(errCode uint32) string {
	if msg, ok := errMessage[errCode]; ok {
		return msg
	}
	return "SERVER INTERNAL ERROR"
}

func IsCodeError(errCode uint32) bool {
	if _, ok := errMessage[errCode]; ok {
		return true
	}
	return false
}
