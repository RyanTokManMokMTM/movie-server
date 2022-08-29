package user

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountFollowingUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountFollowingUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountFollowingUserLogic {
	return &CountFollowingUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountFollowingUserLogic) CountFollowingUser(req *types.CountFollowingReq) (resp *types.CountFollowingResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, req.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	count, err := l.svcCtx.DAO.CountFollowingUser(l.ctx, req.UserId)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	return &types.CountFollowingResp{
		Total: uint(count),
	}, nil
}
