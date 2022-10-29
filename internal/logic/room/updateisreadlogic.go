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

type UpdateIsReadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateIsReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIsReadLogic {
	return &UpdateIsReadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateIsReadLogic) UpdateIsRead(req *types.SetIsReadReq) (resp *types.SetIsReadResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	if userId == 0 {
		return nil, fmt.Errorf("user_id is missing")
	}

	//find that user
	//TODO: Check User is exist
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not exist")
		}
		return nil, err
	}

	_, err = l.svcCtx.DAO.FindOneByRoomID(l.ctx, req.RoomID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("room not exist")
		}
		return nil, err
	}

	if err := l.svcCtx.DAO.UpdateIsRead(l.ctx, req.RoomID, true); err != nil {
		return nil, err
	}
	return &types.SetIsReadResp{
		//IsRead: true,
	}, err
}
