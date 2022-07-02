package listDetail

import (
	"context"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListMoviesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListMoviesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListMoviesLogic {
	return &GetListMoviesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListMoviesLogic) GetListMovies(req *types.ListDetailInfoReq) (resp *types.ListDetailInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
