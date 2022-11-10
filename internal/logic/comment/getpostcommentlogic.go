package comment

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/common/pagination"
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
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//post is exist?
	_, err = l.svcCtx.DAO.FindOnePostInfo(l.ctx, req.PostID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.POST_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	limit := pagination.GetLimit(req.Limit)
	pageOffset := pagination.PageOffset(pagination.DEFAULT_PAGE_SIZE, req.Page)

	commentList, count, err := l.svcCtx.DAO.FindPostComments(l.ctx, req.PostID, int(limit), int(pageOffset))

	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	logx.Info("total record : ", count)

	totalPage := pagination.GetTotalPageByPageSize(uint(count), pagination.DEFAULT_PAGE_SIZE)
	var comments []types.CommentInfo
	for _, v := range commentList {
		comments = append(comments, types.CommentInfo{
			CommentID: v.CommentID,
			UserInfo: types.CommentUser{
				UserID:     v.User.ID,
				UserName:   v.User.Name,
				UserAvatar: v.User.Avatar,
			},
			//PostID:   v.PostID,
			LikesCount:      v.LikesCount,
			Comment:         v.Comment,
			ReplyComment:    uint(len(v.Comments)),
			UpdateAt:        v.UpdatedAt.Unix(),
			ParentCommentID: uint(v.ParentID.Int64),
			IsLiked:         len(v.LikedUser) == 1,
		})
	}

	if len(comments) == 0 {
		comments = make([]types.CommentInfo, 0)
	}
	return &types.GetPostCommentsResp{
		Comments: comments,
		MetaData: types.MetaData{
			TotalPage: totalPage,
		},
	}, nil
}
