package room

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

type GetUserRoomsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRoomsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRoomsLogic {
	return &GetUserRoomsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRoomsLogic) GetUserRooms(req *types.GetUserRoomsReq) (resp *types.GetUserRoomsResp, err error) {
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

	userRooms, err := l.svcCtx.DAO.GetUserRoomsWithMembers(l.ctx, userId)
	if err != nil {
		return nil, err
	}

	roomInfos := make([]types.ChatRoomData, 0)
	for _, v := range userRooms.Rooms {
		total, err := l.svcCtx.DAO.CountMessage(l.ctx, v.ID)
		if err != nil {
			logx.Errorf("count room message error %v ", err)
			continue
		}

		totalPage := pagination.GetTotalPageByPageSize(uint(total), 20)

		user := make([]types.UserInfo, 0)
		for _, u := range v.Users {
			if u.ID == userId {
				continue
			}
			user = append(user, types.UserInfo{
				ID:     u.ID,
				Name:   u.Name,
				Avatar: u.Avatar,
			})
		}

		//The latest 10th messages in asc order
		//get message info

		roomMsg, _ := l.svcCtx.DAO.GetRoomLatestMessage(l.ctx, v.ID)

		messages := make([]types.MessageInfo, 0)
		for i := len(roomMsg) - 1; i >= 0; i-- {
			messages = append(messages, types.MessageInfo{
				ID:              roomMsg[i].MessageID,
				MessageIdentity: roomMsg[i].ID,
				Message:         roomMsg[i].Content,
				Sender:          roomMsg[i].SendUser.ID,
				SentTime:        roomMsg[i].SentTime.Unix(),
			})
		}
		roomInfos = append(roomInfos, types.ChatRoomData{
			ID:           v.ID,
			Users:        user,
			Messages:     messages,
			IsRead:       v.IsRead, //read by other user rather than the sender
			LastSenderID: uint(v.LastSender.Int64),
			MetaData: types.MetaData{
				Page:         1,
				TotalPages:   totalPage,
				TotalResults: uint(total),
			},
		})

	}

	return &types.GetUserRoomsResp{
		Rooms: roomInfos,
	}, nil
}
