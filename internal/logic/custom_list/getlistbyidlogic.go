package custom_list

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errorx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListByIDLogic {
	return &GetListByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListByIDLogic) GetListByID(req *types.UserListReq) (resp *types.UserListResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.List.FindOne(l.ctx, req.ID)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, errorx.NewDefaultCodeError("list not found")
		}
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	return &types.UserListResp{
		List: types.ListInfo{
			ID:       res.ListId,
			Title:    res.ListTitle,
			UpdateOn: res.UpdateTime.Unix(),
		},
	}, nil
}
