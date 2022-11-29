package posts

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

type GetPostByPostIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostByPostIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostByPostIDLogic {
	return &GetPostByPostIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostByPostIDLogic) GetPostByPostID(req *types.PostInfoByIdReq) (resp *types.PostInfoByIdResp, err error) {
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

	postInfo, err := l.svcCtx.DAO.FindOnePostInfoWithUserLiked(l.ctx, req.PostID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.POST_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return &types.PostInfoByIdResp{
		Info: types.PostInfo{
			PostID:           postInfo.PostId,
			PostDesc:         postInfo.PostDesc,
			PostTitle:        postInfo.PostTitle,
			PostCommentCount: int64(len(postInfo.Comments)),
			PostMovie: types.PostMovieInfo{
				MovieID:    postInfo.MovieInfoId,
				Title:      postInfo.MovieInfo.Title,
				PosterPath: postInfo.MovieInfo.PosterPath,
			},
			PostLikeCount: int64(len(postInfo.PostsLiked)),
			PostUser: types.PostUserInfo{
				UserID:     postInfo.UserId,
				UserName:   postInfo.UserInfo.Name,
				UserAvatar: postInfo.UserInfo.Avatar,
			},
			IsPostLikedByUser: len(postInfo.PostsLiked) == 1,
			CreateAt:          postInfo.CreatedAt.Unix(),
		},
	}, nil
}
