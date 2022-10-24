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

	roomInfos := make([]types.RoomInfo, 0)
	for _, v := range userRooms.Rooms {
		roomMembers := make([]types.UserInfo, 0)
		for _, mem := range v.Users {
			roomMembers = append(roomMembers, types.UserInfo{
				ID:     mem.ID,
				Name:   mem.Name,
				Avatar: mem.Avatar,
			})
		}

		roomInfos = append(roomInfos, types.RoomInfo{
			RoomID:   v.ID,
			RoomUser: roomMembers,
		})
	}

	return &types.GetUserRoomsResp{
		Rooms: roomInfos,
	}, nil
}
