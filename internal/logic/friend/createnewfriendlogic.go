package friend

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNewFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateNewFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNewFriendLogic {
	return &CreateNewFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateNewFriendLogic) CreateNewFriend(req *types.CreateNewFriendReq) (resp *types.CreateNewFriendResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	//find that user
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	friend, err := l.svcCtx.DAO.FindOneFriend(l.ctx, userID, req.FriendId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		//create a new record
		if err := l.svcCtx.DAO.CreateNewFriend(l.ctx, userID, req.FriendId); err != nil {
			return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
		}
	} else {
		//just update the state
		friend.State = 1
		if err := l.svcCtx.DAO.UpdateFriendState(l.ctx, friend); err != nil {
			return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
		}
	}

	return &types.CreateNewFriendResp{}, nil
}
