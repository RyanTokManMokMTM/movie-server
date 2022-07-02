package listDetail

import (
	"context"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateListMovieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateListMovieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateListMovieLogic {
	return &CreateListMovieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateListMovieLogic) CreateListMovie(req *types.CreateListDetailInfoReq) (resp *types.CreateListDetailInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
