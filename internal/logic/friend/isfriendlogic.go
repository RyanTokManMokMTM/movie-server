package friend

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFriendLogic {
	return &IsFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsFriendLogic) IsFriend(req *types.IsFriendReq) (resp *types.IsFriendResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	if userId == 0 {
		return nil, fmt.Errorf("user_id is missing")
	}

	//find that user
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not exist")
		}
		return nil, err
	}

	//Check friend
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not exist")
		}
		return nil, err
	}

	//TODO: Get Notification if any
	//TODO: Get Friend if added
	//either send -> receiver or receiver -> sender
	n, err := l.svcCtx.DAO.FindOneFriendNotification(l.ctx, userId, req.UserID)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//Notification not found
			//Check is friend
			isFriend, err := l.svcCtx.DAO.IsFriend(l.ctx, userId, req.UserID)
			if err != nil {
				logx.Error(err)
				return nil, err
			}

			return &types.IsFriendResp{
				IsFriend:      isFriend,
				IsSentRequest: false,
			}, nil

		}
	}

	return &types.IsFriendResp{
		IsFriend:      false,
		IsSentRequest: true,
		RequestInfo: types.BasicRequestInfo{
			RequestID: n.ID,
			SenderID:  n.Sender,
		},
	}, nil
}
