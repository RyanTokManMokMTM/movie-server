package custom_list

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCustomListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCustomListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCustomListLogic {
	return &UpdateCustomListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCustomListLogic) UpdateCustomList(req *types.UpdateCustomListReq) (resp *types.UpdateCustomListResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	//find user
	user, err := l.svcCtx.User.FindOne(l.ctx, userID)
	if err != nil && err != sqlx.ErrNotFound {
		return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UpdateCustomList - user db err:%v, userID:%v", err, userID))
	}

	if user == nil {
		return nil, errors.Wrap(errx.NewErrCode(errx.USER_NOT_EXIST), fmt.Sprintf("UpdateCustomList - user db FINDgot NOT FOUND err: %v, userID: %v", err, userID))
	}

	res, err := l.svcCtx.List.FindOne(l.ctx, req.ID)
	if err != nil {
		return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UpdateCustomList - list db FIND err: %v, ListID: %v", err, req.ID))
	}

	//title is a required field
	res.ListTitle = req.Title

	err = l.svcCtx.List.Update(l.ctx, res)
	if err != nil {
		return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UpdateCustomList - list db UPDATE  err: %v, req: %+v", err, req))
	}
	return &types.UpdateCustomListResp{}, nil
}
