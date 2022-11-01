package comment_likes

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveCommentLikesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveCommentLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveCommentLikesLogic {
	return &RemoveCommentLikesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveCommentLikesLogic) RemoveCommentLikes(req *types.RemoveCommentLikesReq) (resp *types.RemoveCommentLikesResq, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	//find that user
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	comment, err := l.svcCtx.DAO.FindOneComment(l.ctx, req.CommentId)
	if err != nil {
		//Create a new record
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.POST_COMMENT_NOT_EXIST)
		}

		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	commentLikes, err := l.svcCtx.DAO.FindOneCommentLiked(l.ctx, userID, req.CommentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.NOT_LIKE_COMMENT_YET)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	commentLikes.State = 0 //always false

	if err := l.svcCtx.DAO.UpdateCommentLiked(l.ctx, commentLikes, comment); err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return &types.RemoveCommentLikesResq{}, nil
}
