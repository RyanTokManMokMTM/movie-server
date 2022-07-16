package ctxtool

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

var CTXJWTUserID = "user_id"

func GetUserIDFromCTX(ctx context.Context) int64 {
	var userID int64
	if jwtUserID, ok := ctx.Value(CTXJWTUserID).(json.Number); ok {

		if id, err := jwtUserID.Int64(); err == nil {
			userID = id
		} else {
			logx.WithContext(ctx).Info(err)
		}
	}
	return userID
}
