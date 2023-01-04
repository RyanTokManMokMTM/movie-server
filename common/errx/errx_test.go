package errx

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func Test_StatCode(t *testing.T) {
	testCases := []struct {
		Name               string
		Err                *CommonError
		ExpectedStatusCode int
	}{
		{
			Name:               "success",
			Err:                NewErrCode(SUCCESS),
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name:               "Bad Request 1",
			Err:                NewErrCode(REQ_PARAM_ERROR),
			ExpectedStatusCode: http.StatusBadRequest,
		},
		{
			Name:               "Bad Request 2",
			Err:                NewErrCode(USER_PASSWORD_INCORRECT),
			ExpectedStatusCode: http.StatusBadRequest,
		},
		{
			Name:               "Bad Request 3",
			Err:                NewErrCode(USER_UPLOAD_USER_COVER_FAILED),
			ExpectedStatusCode: http.StatusBadRequest,
		},
		{
			Name:               "Unauthorized 1",
			Err:                NewErrCode(TOKEN_EXPIRED_ERROR),
			ExpectedStatusCode: http.StatusUnauthorized,
		},
		{
			Name:               "Unauthorized 2",
			Err:                NewErrCode(TOKEN_INVALID_ERROR),
			ExpectedStatusCode: http.StatusUnauthorized,
		},
		{
			Name:               "InternalError 1",
			Err:                NewErrCode(SERVER_COMMON_ERROR),
			ExpectedStatusCode: http.StatusInternalServerError,
		},
		{
			Name:               "InternalError 2",
			Err:                NewErrCode(UNFOLLOW_FRIEND_ERROR),
			ExpectedStatusCode: http.StatusInternalServerError,
		},
		{
			Name:               "StatusNotFound 1",
			Err:                NewErrCode(USER_NOT_EXIST),
			ExpectedStatusCode: http.StatusNotFound,
		},
		{
			Name:               "StatusNotFound 2",
			Err:                NewErrCode(LIST_MOVIE_NOT_IN_LIST),
			ExpectedStatusCode: http.StatusNotFound,
		},
		{
			Name:               "StatusConflict 1",
			Err:                NewErrCode(EMAIL_HAS_BEEN_REGISTERED),
			ExpectedStatusCode: http.StatusConflict,
		},
		{
			Name:               "StatusConflict 2",
			Err:                NewErrCode(MOVIE_ALREADY_LIKED),
			ExpectedStatusCode: http.StatusConflict,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			code := test.Err.StatusCode()
			require.Equal(t, test.ExpectedStatusCode, code)
		})
	}
}
