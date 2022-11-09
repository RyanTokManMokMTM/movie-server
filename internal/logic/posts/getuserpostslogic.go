package posts

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/common/pagination"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPostsLogic {
	return &GetUserPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPostsLogic) GetUserPosts(req *types.PostsInfoReq) (resp *types.PostsInfoResp, err error) {
	// todo: add your logic here and delete this line
	//userID := ctxtool.GetUserIDFromCTX(l.ctx)

	//find that user
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	limit := pagination.GetLimit(req.Limit)
	pageOffset := pagination.PageOffset(pagination.DEFAULT_PAGE_SIZE, req.Page)

	//Get Post By User ID
	res, count, err := l.svcCtx.DAO.FindUserPosts(l.ctx, req.UserID, int(limit), int(pageOffset))
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	logx.Info("total record : ", count)
	totalPage := pagination.GetTotalPageByPageSize(uint(count), pagination.DEFAULT_PAGE_SIZE)
	//User Post List
	var posts []types.PostInfo
	for _, v := range res {
		var isPostLiked uint = 0
		_, err := l.svcCtx.DAO.FindOnePostLiked(l.ctx, req.UserID, v.PostId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
		}

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			isPostLiked = 1
		}

		posts = append(posts, types.PostInfo{
			PostID:           v.PostId,
			PostDesc:         v.PostDesc,
			PostTitle:        v.PostTitle,
			PostCommentCount: int64(len(v.Comments)),
			PostMovie: types.PostMovieInfo{
				MovieID:    v.MovieInfo.Id,
				Title:      v.MovieInfo.Title,
				PosterPath: v.MovieInfo.PosterPath,
			},
			PostLikeCount: int64(len(v.PostsLiked)),
			PostUser: types.PostUserInfo{
				UserID:     v.UserInfo.ID,
				UserName:   v.UserInfo.Name,
				UserAvatar: v.UserInfo.Avatar,
			},
			IsPostLikedByUser: isPostLiked != 0,
			CreateAt:          v.CreatedAt.Unix(),
		})
	}
	return &types.PostsInfoResp{
		Infos: posts,
		MetaData: types.MetaData{
			TotalPage: totalPage,
		},
	}, nil
}
