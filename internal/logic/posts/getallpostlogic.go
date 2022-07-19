package posts

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"log"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPostLogic {
	return &GetAllPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllPostLogic) GetAllPost(req *types.PostsInfoReq) (resp *types.PostsInfoResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.PostModel.FindAllWithInfoByCreateTime(l.ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	var infos []types.PostInfo
	for _, v := range res {
		infos = append(infos, types.PostInfo{
			PostID:           v.PostId,
			PostTitle:        v.PostTitle,
			PostDesc:         v.PostDesc,
			PostLikeCount:    v.PostLike,
			PostCommentCount: v.CommentCount,
			CreateAt:         v.CreateTime.Unix(),
			//UpdateTime:       v.UpdateTime.Unix(),
			PostMovie: types.PostMovieInfo{
				MovieID:    v.MovieId,
				Title:      v.MovieTitle,
				PosterPath: v.MoviePoster,
			},
			PostUser: types.PostUserInfo{
				UserID:     v.UserId,
				UserName:   v.UserName,
				UserAvatar: v.UserAvatar,
			},
		})
	}

	return &types.PostsInfoResp{
		Infos: infos,
	}, nil
}
