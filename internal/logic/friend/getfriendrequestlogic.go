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

type GetFriendRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendRequestLogic {
	return &GetFriendRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendRequestLogic) GetFriendRequest(req *types.GetFriendRequestReq) (resp *types.GetFriendRequestResp, err error) {
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

	list, err := l.svcCtx.DAO.GetFriendRequest(l.ctx, userId)
	if err != nil {
		return nil, err
	}

	requests := make([]types.FriendRequest, 0)
	for _, req := range list {
		requests = append(requests, types.FriendRequest{
			RequestID: req.ID,
			Sender: types.UserInfo{
				ID:     req.SenderInfo.ID,
				Name:   req.SenderInfo.Name,
				Avatar: req.SenderInfo.Avatar,
			},
			SentTime: req.CreatedAt.Unix(),
			State:    req.State,
		})
	}
	return &types.GetFriendRequestResp{
		Requests: requests,
	}, nil
}
