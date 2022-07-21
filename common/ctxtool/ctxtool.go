package ctxtool

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

var CTXJWTUserID = "user_id"

func GetUserIDFromCTX(ctx context.Context) uint {
	var userID uint
	if jwtUserID, ok := ctx.Value(CTXJWTUserID).(json.Number); ok {

		if id, err := jwtUserID.Int64(); err == nil {
			userID = uint(id)
		} else {
			logx.WithContext(ctx).Info(err)
		}
	}
	return userID
}
