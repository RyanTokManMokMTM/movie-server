package comment_likes

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountCommentLikesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountCommentLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountCommentLikesLogic {
	return &CountCommentLikesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountCommentLikesLogic) CountCommentLikes(req *types.CountCommentLikesReq) (resp *types.CountCommentLikesResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.DAO.FindOneComment(l.ctx, req.CommentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.POST_COMMENT_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	total, err := l.svcCtx.DAO.CountCommentLikes(l.ctx, req.CommentId)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	return &types.CountCommentLikesResp{
		TotalLikes: uint(total),
	}, nil
}
