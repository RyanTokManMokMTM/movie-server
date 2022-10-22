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

type RemoveFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveFriendLogic {
	return &RemoveFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveFriendLogic) RemoveFriend(req *types.RemoveFriendReq) (resp *types.RemoveFriendResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	if userId == 0 {
		return nil, fmt.Errorf("user_id is missing")
	}

	//find that user.
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not exist")
		}
		return nil, err
	}

	//TODO: Has a friendship?
	f, err := l.svcCtx.DAO.HasFriendShip(l.ctx, userId, req.FriendID)
	if err != nil {
		return nil, err
	}

	if f.ID == 0 {
		return nil, fmt.Errorf("both of you are not friend")
	}

	//TODO: Remove the friendship
	if err := l.svcCtx.DAO.RemoveFriend(l.ctx, userId, req.FriendID); err != nil {
		return nil, err
	}

	return &types.RemoveFriendResp{}, nil
}
