package notification

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/common/pagination"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetlikenotificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetlikenotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetlikenotificationLogic {
	return &GetlikenotificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetlikenotificationLogic) Getlikenotification(req *types.GetLikeNotificationReq) (resp *types.GetLikeNotificationResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	limit := pagination.GetLimit(req.Limit)
	pageOffset := pagination.PageOffset(pagination.DEFAULT_PAGE_SIZE, req.Page)

	list, count, err := l.svcCtx.DAO.FindLikesNotificationByReceiver(l.ctx, userID, int(limit), int(pageOffset))
	if err != nil {
		return nil, err
	}
	logx.Info("total record : ", count)

	totalPage := pagination.GetTotalPageByPageSize(uint(count), pagination.DEFAULT_PAGE_SIZE)
	res := make([]types.LikedNotification, 0)
	for _, v := range list {
		res = append(res, types.LikedNotification{
			ID: v.ID,
			LikedBy: types.UserInfo{
				ID:     v.LikedUser.ID,
				Name:   v.LikedUser.Name,
				Avatar: v.LikedUser.Avatar,
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
			Type:    v.Type,
			LikedAt: uint(v.LikedTime.Unix()),
		})
	}

	return &types.GetLikeNotificationResp{
		LikedNotificationList: res,
		MetaData: types.MetaData{
			TotalPage: totalPage,
		},
	}, nil
}
