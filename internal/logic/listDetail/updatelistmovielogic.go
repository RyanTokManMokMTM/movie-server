package listDetail

import (
	"context"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateListMovieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateListMovieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateListMovieLogic {
	return &UpdateListMovieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateListMovieLogic) UpdateListMovie(req *types.UpdateListDetailInfoReq) (resp *types.UpdateListDetailInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
