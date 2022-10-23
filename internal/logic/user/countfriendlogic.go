package user

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountFriendLogic {
	return &CountFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountFriendLogic) CountFriend(req *types.CountFriendReq) (resp *types.CountFriendResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, req.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//TODO: Get User Friend
	//f, err := l.svcCtx.DAO.GetUserFriendRecord(l.ctx, req.UserId)
	//if err != nil {
	//	return nil, err
	//}

	count := l.svcCtx.DAO.CountFriends(l.ctx, req.UserId)
	//if err != nil {
	//	return nil, err
	//}
	return &types.CountFriendResp{
		Total: uint(count),
	}, nil
}
