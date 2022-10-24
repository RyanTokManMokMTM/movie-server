package message

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

type GetRoomMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoomMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomMessageLogic {
	return &GetRoomMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoomMessageLogic) GetRoomMessage(req *types.GetRoomMessageReq) (resp *types.GetRoomMessageResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	if userId == 0 {
		return nil, fmt.Errorf("user_id is missing")
	}

	//find that user
	u, err := l.svcCtx.DAO.FindUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not exist")
		}
		return nil, err
	}

	//TODO: Check User is joined the group
	mem, err := l.svcCtx.DAO.FindOneRoomMember(l.ctx, req.RoomID, u.ID)
	if err != nil {
		return nil, err
	}

	if mem.ID == 0 {
		return nil, fmt.Errorf("you haven't joined the group")
	}

	//TODO: Get at most 10 latest record belong to the group
	msgs, err := l.svcCtx.DAO.GetRoomMessage(l.ctx, req.RoomID)
	if err != nil {
		return nil, err
	}

	record := make([]types.MessageData,0)
	for _, data := range msgs {
		record = append(record, types.MessageData{
			MessageID: data.ID,
			UserInfo: types.UserInfo{
				ID:     data.SendUser.ID,
				Name:   data.SendUser.Name,
				Avatar: data.SendUser.Avatar,
			},
			Content:  data.Data,
			SendTime: data.CreatedAt.Unix(),
		})
	}

	return &types.GetRoomMessageResp{
		Messagees: record,
	}, nil
}
