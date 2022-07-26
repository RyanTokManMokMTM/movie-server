package posts

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
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
	//find that user
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//Get Post By User ID
	res, err := l.svcCtx.DAO.FindUserPosts(l.ctx, req.UserID)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//User Post List
	var posts []types.PostInfo
	for _, v := range res {
		posts = append(posts, types.PostInfo{
			PostID:           v.PostId,
			PostDesc:         v.PostDesc,
			PostTitle:        v.PostTitle,
			PostCommentCount: int64(len(v.Comments)),
			PostMovie: types.PostMovieInfo{
				MovieID:    v.MovieInfo.MovieId,
				Title:      v.MovieInfo.Title,
				PosterPath: v.MovieInfo.PosterPath,
			},
			PostLikeCount: v.PostLike,
			PostUser: types.PostUserInfo{
				UserID:     v.UserInfo.Id,
				UserName:   v.UserInfo.Name,
				UserAvatar: v.UserInfo.Avatar,
			},
			CreateAt: v.CreatedAt.Unix(),
		})
	}
	return &types.PostsInfoResp{
		Infos: posts,
	}, nil
}
