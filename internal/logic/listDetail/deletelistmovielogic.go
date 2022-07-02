package listDetail

import (
	"context"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteListMovieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteListMovieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteListMovieLogic {
	return &DeleteListMovieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteListMovieLogic) DeleteListMovie(req *types.DeleteListDetailInfoReq) (resp *types.DeleteListDetailInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
