package user

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendListLogic) GetFriendList(req *types.GetFriendListReq) (resp *types.GetFriendListResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, req.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//TODO: User Friend Record
	f, err := l.svcCtx.DAO.GetUserFriendRecord(l.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	list, err := l.svcCtx.DAO.GetFriendsList(l.ctx, f.ID)
	if err != nil {
		return nil, err
	}

	var friendList []types.UserInfo
	for _, info := range list {
		friendList = append(friendList, types.UserInfo{
			ID:     info.ID,
			Name:   info.Name,
			Avatar: info.Avatar,
		})
	}
	return &types.GetFriendListResp{
		Friends: friendList,
	}, nil
}
