package errx

import "github.com/zeromicro/go-zero/core/logx"

var errMessage map[uint32]string

func init() {
	errMessage = make(map[uint32]string)

	errMessage[SUCCESS] = "SUCCESS"
	errMessage[SERVER_COMMON_ERROR] = "SERVER INTERNAL ERROR"
	errMessage[REQ_PARAM_ERROR] = "REQUEST PARAMETER ERROR"
	errMessage[TOKEN_EXPIRED_ERROR] = "TOKEN HAS BEEN EXPIRED"
	errMessage[TOKEN_INVALID_ERROR] = "TOKEN HAS BEEN Invalid"
	errMessage[TOKEN_GENERATE_ERROR] = "TOKEN GENERATE FAILED"
	errMessage[DB_ERROR] = "DATABASE ERROR"
	errMessage[DB_AFFECTED_ZERO_ERROR] = "DATABASE AFFECTED 0 rows"
	errMessage[EMAIL_HAS_BEEN_REGISTERED] = "USER HAS BEEN REGISTERED"
	errMessage[USER_NOT_EXIST] = "USER NOT EXIST"
	errMessage[USER_PASSWORD_INCORRECT] = "USER PASSWORD INCORRECT"
	errMessage[USER_UPLOAD_USER_AVATAR_FAILED] = "USER UPLOAD AVATAR FAILED"
	errMessage[USER_UPLOAD_USER_COVER_FAILED] = "USER UPLOAD COVER FAILED"

	errMessage[MOVIE_NOT_EXIST] = "MOVIE NOT EXIST"
	errMessage[MOVIE_ALREADY_LIKED] = "MOVIE IS ALREADY LIKED"

	errMessage[POST_NOT_EXIST] = "POST NOT EXIST"
	errMessage[LIST_NOT_EXIST] = "LIST NOT EXIST"
	errMessage[LIST_MOVIE_ALREADY_IN_LIST] = "MOVIE ALREADY IN LIST"
	errMessage[LIST_MOVIE_NOT_IN_LIST] = "MOVIE NOT IN LIST"
	errMessage[POST_COMMENT_NOT_EXIST] = "COMMENT NOT EXIST"
	errMessage[FOLLOW_FRIEND_ERROR] = "FAILED TO FOLLOW USER"
	errMessage[UNFOLLOW_FRIEND_ERROR] = "FAILED TO UNFOLLOW USER"

	errMessage[WEBSOCKET_CONNECTION_ERROR] = "UPGRADE TO WEB SOCKET FAILED"
	errMessage[WEBSOCKET_READ_MESSAG_ERROR] = "READ DATA FROM WEBSOCKET FAILED"
	errMessage[WEBSOCKET_WRITE_MESSAGE_ERROR] = "WRITE DATA TO WEBSOCKET FAILED"

}

func MapErrMsg(errCode uint32) string {
	logx.Info(errCode)
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
