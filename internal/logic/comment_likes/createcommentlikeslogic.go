package comment_likes

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/logic/serverWs"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"gorm.io/gorm"
	"time"

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
	u, err := l.svcCtx.DAO.FindUserByID(l.ctx, userID)
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

	//TODO: add comment liked
	logx.Infof("%+v", comment)
	if err := l.svcCtx.DAO.InsertOneCommentLike(l.ctx, userID, comment.CommentID, comment.LikesCount+1); err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//TODO: Send notification
	if userID != comment.UserID {
		//TODO: is notification exist?
		err = l.svcCtx.DAO.FindOneLikeCommentNotification(l.ctx, comment.UserID, userID, comment.CommentID)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			logx.Info("Notification not found...")

			if err := l.svcCtx.DAO.InsertOneCommentLikeNotification(l.ctx, comment.PostID, userID, req.CommentId, comment.UserID, time.Now()); err != nil {
				return nil, err
			}

			go func() {
				_ = serverWs.SendNotificationToUserWithUserInfo(comment.UserID, u, fmt.Sprintf("%s給您的留言點讚", u.Name))
			}()

		}
	}

	return &types.CreateCommentLikesResp{}, nil
}
