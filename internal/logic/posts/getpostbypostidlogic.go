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

func (l *GetPostByPostIDLogic) GetPostByPostID(req *types.PostInfoReq) (resp *types.PostInfoResp, err error) {
	// todo: add your logic here and delete this line
	postInfo, err := l.svcCtx.DAO.GetPostInfo(l.ctx, req.PostID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.POST_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return &types.PostInfoResp{
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
			PostLikeCount: postInfo.PostLike,
			PostUser: types.PostUserInfo{
				UserID:     postInfo.UserId,
				UserName:   postInfo.UserInfo.Name,
				UserAvatar: postInfo.UserInfo.Avatar,
			},
			CreateAt: postInfo.CreatedAt.Unix(),
		},
	}, nil
}
