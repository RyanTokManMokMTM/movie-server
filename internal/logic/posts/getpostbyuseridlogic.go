package posts

import (
	"context"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostByUserIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostByUserIDLogic {
	return &GetPostByUserIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostByUserIDLogic) GetPostByUserID(req *types.PostInfoByUserReq) (resp *types.PostInfoByUserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
