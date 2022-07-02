package UserMovieList

import (
	"context"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllListLogic {
	return &GetAllListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllListLogic) GetAllList(req *types.ListsReq) (resp *types.ListsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
