package user

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFollowingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserFollowingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFollowingListLogic {
	return &GetUserFollowingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserFollowingListLogic) GetUserFollowingList(req *types.GetFollowingListReq) (resp *types.GetFollowingListResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, req.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	followingList, err := l.svcCtx.DAO.FindUserFollowingList(l.ctx, req.UserId)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	var users []types.UserInfo
	for _, user := range followingList {
		users = append(users, types.UserInfo{
			ID:     user.Id,
			Name:   user.Name,
			Avatar: user.Avatar,
		})
	}
	return &types.GetFollowingListResp{
		Users: users,
	}, nil
}
