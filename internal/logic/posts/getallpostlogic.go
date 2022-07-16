package posts

import (
	"context"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPostLogic {
	return &GetAllPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllPostLogic) GetAllPost(req *types.PostsInfoReq) (resp *types.PostsInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
