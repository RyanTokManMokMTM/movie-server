package room

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

type GetRoomInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoomInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomInfoLogic {
	return &GetRoomInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoomInfoLogic) GetRoomInfo(req *types.GetRoomInfoReq) (resp *types.GetRoomInfoResp, err error) {
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

	//TODO: RoomExist?

	room, err := l.svcCtx.DAO.FindOneByRoomIDWithInfo(l.ctx, req.RoomID)
	if err != nil {
		return nil, err
	}

	//The latest 10th messages in asc order
	messages := make([]types.MessageInfo, 0)
	user := make([]types.UserInfo, 0)

	//just return the last message instead(id and message?)
	for i := len(room.Messages) - 1; i >= 0; i-- {
		messages = append(messages, types.MessageInfo{
			ID:       room.Messages[i].MessageID,
			Message:  room.Messages[i].Content,
			Sender:   room.Messages[i].SendUser.ID,
			SentTime: room.Messages[i].SentTime.Unix(),
		})
	}

	for _, u := range room.Users {
		if u.ID == userId {
			continue
		}

		user = append(user, types.UserInfo{
			ID:     u.ID,
			Name:   u.Name,
			Avatar: u.Avatar,
		})
	}

	roomInfos := types.ChatRoomData{
		ID:           room.ID,
		Users:        user,
		Messages:     messages,
		IsRead:       room.IsRead, //read by other user rather than the sender
		LastSenderID: uint(room.LastSender.Int64),
	}

	return &types.GetRoomInfoResp{
		Info: roomInfos,
	}, nil
}
