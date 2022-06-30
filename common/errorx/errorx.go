package errorx

var defaultCode = 1001

type CodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CodeErrorResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewCodeError(code int, message string) error {
	return &CodeError{
		Code:    code,
		Message: message,
	}
}

func NewDefaultCodeError(message string) error {
	return &CodeError{
		Code:    defaultCode,
		Message: message,
	}
}

func (err *CodeError) Error() string {
	return err.Message
}

func (err *CodeError) DataResponse() *CodeErrorResp {
	return &CodeErrorResp{
		Code:    err.Code,
		Message: err.Message,
	}
}
