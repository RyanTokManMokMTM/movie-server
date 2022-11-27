package comment

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/logic/serverWs"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateReplyCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateReplyCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateReplyCommentLogic {
	return &CreateReplyCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateReplyCommentLogic) CreateReplyComment(req *types.CreateReplyCommentReq) (resp *types.CreateReplyCommentResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	u, err := l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//check post is exist
	post, err := l.svcCtx.DAO.FindOnePostInfo(l.ctx, req.PostID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.POST_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//Comment is exist?
	replyComment, err := l.svcCtx.DAO.FindOneComment(l.ctx, req.ReplyCommentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.POST_COMMENT_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	comment, err := l.svcCtx.DAO.CreatePostReplyComment(l.ctx, userID, post.PostId, req.ReplyCommentId, req.ParentCommentID, replyComment.UserID, req.Comment)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//TODO:create a reply comment notification
	//TODO:- no need to notify comment owner - Owner -> Owner
	if userID != replyComment.UserID {
		if err := l.svcCtx.DAO.InsertOneReplyCommentNotification(l.ctx, replyComment.UserID, userID, post.PostId, comment.CommentID, replyComment.CommentID, comment.CreatedAt); err != nil {
			return nil, err
		}
		//send a notification
		go func() {
			logx.Info("send a reply comment notification")
			_ = serverWs.SendNotificationToUserWithUserInfo(replyComment.UserID, u, fmt.Sprintf("%s回覆您的留言", u.Name), serverWs.COMMENT_NOTIFICATION)
		}()
	}

	return &types.CreateReplyCommentResp{
		CommentID: comment.CommentID,
		CreateAt:  comment.CreatedAt.Unix(),
	}, nil
}
