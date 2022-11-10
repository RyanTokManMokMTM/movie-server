package friend

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/pagination"
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

	limit := pagination.GetLimit(req.Limit)
	pageOffset := pagination.PageOffset(pagination.DEFAULT_PAGE_SIZE, req.Page)

	list, count, err := l.svcCtx.DAO.GetFriendRequest(l.ctx, userId, int(limit), int(pageOffset))
	if err != nil {
		return nil, err
	}
	logx.Info("total record : ", count)

	totalPage := pagination.GetTotalPageByPageSize(uint(count), pagination.DEFAULT_PAGE_SIZE)
	requests := make([]types.FriendRequest, 0)
	for _, req := range list {
		//if user is sender ,send receiver data
		//else send sender data
		var info types.UserInfo
		if userId == req.SenderInfo.ID { //if user is sender -> it needs the receiver info
			info = types.UserInfo{
				ID:     req.ReceiverInfo.ID,
				Name:   req.ReceiverInfo.Name,
				Avatar: req.ReceiverInfo.Avatar,
			}
		} else {
			info = types.UserInfo{
				ID:     req.SenderInfo.ID,
				Name:   req.SenderInfo.Name,
				Avatar: req.SenderInfo.Avatar,
			}
		}

		requests = append(requests, types.FriendRequest{
			RequestID: req.ID,
			Sender:    info,
			SentTime:  req.CreatedAt.Unix(),
			State:     req.State,
		})
	}
	return &types.GetFriendRequestResp{
		Requests: requests,
		MetaData: types.MetaData{
			TotalPages:   totalPage,
			TotalResults: uint(count),
			Page:         pagination.GetPage(req.Page),
		},
	}, nil
}
