package errx

import (
	"fmt"
	"net/http"
)

type CommonError struct {
	code InternalCode
	msg  string
}

type CommonErrorResp struct {
	Code InternalCode `json:"code"`
	Msg  string       `json:"message"`
}

func (e *CommonError) Error() string {
	return fmt.Sprintf("Code: %v, Msg:%v", e.code, e.msg)
}

func (e *CommonError) GetCode() InternalCode {
	return e.code
}

func (e *CommonError) GetErrorMessage() string {
	return e.msg
}

func (e *CommonError) StatusCode() int {
	switch e.code {
	case SUCCESS:
		return http.StatusOK
	case REQ_PARAM_ERROR:
		fallthrough
	case USER_PASSWORD_INCORRECT:
		fallthrough
	case USER_UPLOAD_USER_AVATAR_FAILED:
		fallthrough
	case USER_UPLOAD_USER_COVER_FAILED:
		return http.StatusBadRequest

	case TOKEN_EXPIRED_ERROR:
		fallthrough
	case TOKEN_INVALID_ERROR:
		return http.StatusUnauthorized

	case SERVER_COMMON_ERROR:
		fallthrough
	case TOKEN_GENERATE_ERROR:
		fallthrough
	case DB_ERROR:
		fallthrough
	case DB_AFFECTED_ZERO_ERROR:
		fallthrough
	case FOLLOW_FRIEND_ERROR:
		fallthrough
	case UNFOLLOW_FRIEND_ERROR:
		fallthrough
	case WEBSOCKET_CONNECTION_ERROR:
		fallthrough
	case WEBSOCKET_READ_MESSAG_ERROR:
		fallthrough
	case WEBSOCKET_WRITE_MESSAGE_ERROR:
		return http.StatusInternalServerError

	case USER_NOT_EXIST:
		fallthrough
	case MOVIE_NOT_EXIST:
		fallthrough
	case LIST_NOT_EXIST:
		fallthrough
	case POST_NOT_EXIST:
		fallthrough
	case POST_COMMENT_NOT_EXIST:
		fallthrough
	case NOT_LIKE_POST_YET:
		fallthrough
	case NOT_LIKE_COMMENT_YET:
		fallthrough
	case MOVIE_NOT_YET_LIKED:
		fallthrough
	case LIST_MOVIE_NOT_IN_LIST:
		return http.StatusNotFound

	case EMAIL_HAS_BEEN_REGISTERED:
		fallthrough
	case MOVIE_ALREADY_LIKED:
		fallthrough
	case LIST_MOVIE_ALREADY_IN_LIST:
		return http.StatusConflict
	}
	return http.StatusInternalServerError
}

func NewCommonMessage(code InternalCode, errMsg string) *CommonError {
	return &CommonError{
		code: code,
		msg:  errMsg,
	}
}

func NewErrMsg(errMsg string) *CommonError {
	return &CommonError{code: SERVER_COMMON_ERROR, msg: errMsg}
}

func NewErrCode(errCode InternalCode) *CommonError {
	return &CommonError{code: errCode, msg: MapErrMsg(errCode)}
}

func (e *CommonError) ToJSONResp() *CommonErrorResp {
	return &CommonErrorResp{
		Code: e.code,
		Msg:  e.msg,
	}
}
