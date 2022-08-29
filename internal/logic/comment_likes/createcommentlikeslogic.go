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

type CreateCommentLikesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLikesLogic {
	return &CreateCommentLikesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLikesLogic) CreateCommentLikes(req *types.CreateCommentLikesReq) (resp *types.CreateCommentLikesResp, err error) {
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

	commentLiked, err := l.svcCtx.DAO.FindOneCommentLiked(l.ctx, userID, req.CommentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//Create a new record
			if err := l.svcCtx.DAO.CreateCommentLiked(l.ctx, userID, comment); err != nil {
				return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
			}

			return &types.CreateCommentLikesResp{}, nil
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	commentLiked.State = 1

	if err := l.svcCtx.DAO.UpdateCommentLiked(l.ctx, commentLiked); err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return &types.CreateCommentLikesResp{}, nil
}
