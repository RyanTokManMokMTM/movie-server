package errx

import "fmt"

type CommonError struct {
	code uint32
	msg  string
}

type CommonErrorResp struct {
	Code uint32 `json:"code"`
	Msg  string `json:"message"`
}

func (e *CommonError) Error() string {
	return fmt.Sprintf("Code: %v, Msg:%v", e.code, e.msg)
}

func (e *CommonError) GetCode() uint32 {
	return e.code
}

func (e *CommonError) GetErrorMessage() string {
	return e.msg
}

func NewCommonMessage(code uint32, errMsg string) *CommonError {
	return &CommonError{
		code: code,
		msg:  errMsg,
	}
}

func NewErrMsg(errMsg string) *CommonError {
	return &CommonError{code: SERVER_COMMON_ERROR, msg: errMsg}
}

func NewErrCode(errCode uint32) *CommonError {
	return &CommonError{code: errCode, msg: MapErrMsg(errCode)}
}

func (e *CommonError) ToJSONResp() *CommonErrorResp {
	return &CommonErrorResp{
		Code: e.code,
		Msg:  e.msg,
	}
}
