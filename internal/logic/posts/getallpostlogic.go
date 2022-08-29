package posts

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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

func (l *GetAllPostLogic) GetAllPost(req *types.AllPostsInfoReq) (resp *types.AllPostsInfoResp, err error) {
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

	res, err := l.svcCtx.DAO.FindAllPosts(l.ctx, userID)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//Post List
	var posts []types.PostInfo
	for _, v := range res {

		var isPostLiked uint = 0
		_, err := l.svcCtx.DAO.FindOnePostLiked(l.ctx, userID, v.PostId)
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
				UserID:     v.UserInfo.Id,
				UserName:   v.UserInfo.Name,
				UserAvatar: v.UserInfo.Avatar,
			},
			CreateAt:          v.CreatedAt.Unix(),
			IsPostLikedByUser: isPostLiked,
		})
	}

	return &types.AllPostsInfoResp{
		Infos: posts,
	}, nil
}

//
//func (l *GetAllPostLogic) GetAllPost(req *types.PostsInfoReq) (resp *types.PostsInfoResp, err error) {
//	// todo: add your logic here and delete this line
//	//res, err := l.svcCtx.PostModel.FindAllWithInfoByCreateTime(l.ctx)
//	//if err != nil {
//	//	log.Println(err.Error())
//	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	//}
//	//
//	//var infos []types.PostInfo
//	//for _, v := range res {
//	//	infos = append(infos, types.PostInfo{
//	//		PostID:           v.PostId,
//	//		PostTitle:        v.PostTitle,
//	//		PostDesc:         v.PostDesc,
//	//		PostLikeCount:    v.PostLike,
//	//		PostCommentCount: v.CommentCount,
//	//		CreateAt:         v.CreateTime.Unix(),
//	//		//UpdateTime:       v.UpdateTime.Unix(),
//	//		PostMovie: types.PostMovieInfo{
//	//			MovieID:    v.Id,
//	//			Title:      v.MovieTitle,
//	//			PosterPath: v.MoviePoster,
//	//		},
//	//		PostUser: types.PostUserInfo{
//	//			UserID:     v.UserId,
//	//			UserName:   v.UserName,
//	//			UserAvatar: v.UserAvatar,
//	//		},
//	//	})
//	//}
//	//
//	//return &types.PostsInfoResp{
//	//	Infos: infos,
//	//}, nil
//	return
//}
