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

type CreateRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoomLogic {
	return &CreateRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoomLogic) CreateRoom(req *types.CreateRoomReq) (resp *types.CreateRoomResp, err error) {
	// todo: add your logic here and delete this line
	//TODO: Get UserID
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

	//TODO: Create A new Room
	r, err := l.svcCtx.DAO.InsertOneRoom(l.ctx, req.Name, req.Info, u.ID)
	if err != nil {
		return nil, err
	}

	//join the room too
	if err := l.svcCtx.DAO.JoinOneRoom(l.ctx, r.ID, u); err != nil {
		return nil, err
	}

	return &types.CreateRoomResp{
		RoomID: r.ID,
		Name:   r.Name,
		Info:   r.Info,
	}, nil
}
