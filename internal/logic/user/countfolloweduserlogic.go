package user

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountFollowedUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountFollowedUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountFollowedUserLogic {
	return &CountFollowedUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountFollowedUserLogic) CountFollowedUser(req *types.CountFollowedReq) (resp *types.CountFollowedResp, err error) {
	// todo: add your logic here and delete this line
	count, err := l.svcCtx.DAO.CountFollowedUser(l.ctx, req.UserId)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	return &types.CountFollowedResp{
		Total: uint(count),
	}, nil
}
