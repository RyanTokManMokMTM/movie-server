package listDetail

import (
	"context"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteListMvoieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteListMvoieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteListMvoieLogic {
	return &DeleteListMvoieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteListMvoieLogic) DeleteListMvoie(req *types.DeleteListDetailInfoReq) (resp *types.DeleteListDetailInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
