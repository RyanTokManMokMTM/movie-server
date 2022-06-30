package errorx

var (
	Success = NewCodeError(0, "succeed")

	ServerError              = NewCodeError(100, "server internal error")
	InvalidParams            = NewCodeError(101, "parameters invalid")
	NotFound                 = NewCodeError(102, "not found")
	UnauthorizedAuthNotExist = NewCodeError(103, "authorized failed, required key is not exist")

	UnauthorizedTokenError         = NewCodeError(104, "authorized failed, token error")
	UnauthorizedTokenTimeOut       = NewCodeError(105, "authorized failed, token expired")
	UnauthorizedTokenGenerateError = NewCodeError(106, "authorized failed, token generation failed")
	TooManyRequest                 = NewCodeError(107, "request is overloaded")
)
