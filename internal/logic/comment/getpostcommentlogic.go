package comment

import (
	"context"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostCommentLogic {
	return &GetPostCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostCommentLogic) GetPostComment(req *types.GetPostCommentsReq) (resp *types.GetPostCommentsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
