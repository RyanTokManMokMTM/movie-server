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

	//post is exist?
	_, err = l.svcCtx.DAO.FindOnePostInfo(l.ctx, req.PostID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.POST_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	commentList, err := l.svcCtx.DAO.FindPostComments(l.ctx, req.PostID)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	var comments []types.CommentInfo
	for _, v := range commentList {
		comments = append(comments, types.CommentInfo{
			CommentID: v.CommentID,
			UserInfo: types.CommentUser{
				UserID:     v.User.Id,
				UserName:   v.User.Name,
				UserAvatar: v.User.Avatar,
			},
			//PostID:   v.PostID,
			Comment:  v.Comment,
			UpdateAt: v.UpdatedAt.Unix(),
		})
	}
	return &types.GetPostCommentsResp{
		Comments: comments,
	}, nil
}
