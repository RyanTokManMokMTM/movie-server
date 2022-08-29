package comment

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReplyCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetReplyCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReplyCommentLogic {
	return &GetReplyCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetReplyCommentLogic) GetReplyComment(req *types.GetReplyCommentReq) (resp *types.GetReplyCommentResp, err error) {
	// todo: add your logic here and delete this line

	//is comment exist
	_, err = l.svcCtx.DAO.FindOneComment(l.ctx, req.CommentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.POST_COMMENT_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//get reply comments list
	var replyComments []types.CommentInfo
	replyList, err := l.svcCtx.DAO.FindReplyComments(l.ctx, req.CommentId)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	for _, reply := range replyList {
		replyComments = append(replyComments, types.CommentInfo{
			CommentID: reply.CommentID,
			UserInfo: types.CommentUser{
				UserID:     reply.User.Id,
				UserName:   reply.User.Name,
				UserAvatar: reply.User.Avatar,
			},
			Comment:      reply.Comment,
			UpdateAt:     reply.UpdatedAt.Unix(),
			ReplyComment: 0, //the parent of this comment is reply post id , it may add a `reply user field` for identifying  which user is replying to. So it is always zero.
		})
	}

	return &types.GetReplyCommentResp{
		ReplyComments: replyComments,
	}, nil
}