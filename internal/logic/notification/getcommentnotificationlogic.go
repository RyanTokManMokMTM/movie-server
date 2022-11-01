package notification

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetcommentnotificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetcommentnotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetcommentnotificationLogic {
	return &GetcommentnotificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetcommentnotificationLogic) Getcommentnotification(req *types.GetCommentNotificationReq) (resp *types.GetCommentNotificationResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	list, err := l.svcCtx.DAO.FindOneCommentNotification(l.ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]types.CommentNotification, 0)
	for _, v := range list {
		res = append(res, types.CommentNotification{
			ID: v.ID,
			CommentBy: types.UserInfo{
				ID:     v.CommentUser.ID,
				Name:   v.CommentUser.Name,
				Avatar: v.CommentUser.Avatar,
			},
			PostInfo: types.SimplePostInfo{
				PostID: v.PostInfo.PostId,
				PostMovie: types.PostMovieInfo{
					MovieID:    v.PostInfo.MovieInfo.Id,
					PosterPath: v.PostInfo.MovieInfo.PosterPath,
					Title:      v.PostInfo.MovieInfo.Title,
				},
			},
			CommentInfo: types.SimpleCommentInfo{
				CommentID: v.CommentInfo.CommentID,
				Comment:   v.CommentInfo.Comment,
			},
			ReplyCommentInfo: types.SimpleCommentInfo{
				CommentID: v.RelyCommentInfo.CommentID,
				Comment:   v.RelyCommentInfo.Comment,
			},
			CommentAt: uint(v.CreatedAt.Unix()),
			Type:      v.Type,
		})
	}

	return &types.GetCommentNotificationResp{
		CommentNotificationList: res,
	}, nil
}
