package user

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	//id := l.ctx.Value("user_id").(json.Number)
	res, err := l.svcCtx.User.FindOne(l.ctx, req.ID)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, errors.New("user not exists")
		}
		return nil, errors.New(err.Error())
	}
	return &types.UserInfoResponse{
		ID:    res.Id,
		Email: res.Email,
		Name:  res.Name,
	}, nil
}
