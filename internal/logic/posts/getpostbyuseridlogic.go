package posts

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostByUserIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostByUserIDLogic {
	return &GetPostByUserIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostByUserIDLogic) GetPostByUserID(req *types.PostInfosByUserReq) (resp *types.PostInfosByUserResp, err error) {
	// todo: add your logic here and delete this line

	//user, err := l.svcCtx.User.FindOne(l.ctx, req.UserID)
	//if err != nil && err != sqlx.ErrNotFound {
	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	//}
	//
	//if user == nil {
	//	return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
	//}
	//
	//res, err := l.svcCtx.PostModel.FindUserWithInfoByCreateTime(l.ctx, req.UserID)
	//if err != nil {
	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	//}
	//
	//var infos []types.PostInfo
	//for _, v := range res {
	//	infos = append(infos, types.PostInfo{
	//		PostID:           v.PostId,
	//		PostTitle:        v.PostTitle,
	//		PostDesc:         v.PostDesc,
	//		PostLikeCount:    v.PostLike,
	//		PostCommentCount: v.CommentCount,
	//		CreateAt:         v.CreateTime.Unix(),
	//		PostMovie: types.PostMovieInfo{
	//			MovieID:    v.MovieId,
	//			Title:      v.MovieTitle,
	//			PosterPath: v.MoviePoster,
	//		},
	//
	//		PostUser: types.PostUserInfo{
	//			UserID:     v.UserId,
	//			UserName:   v.UserName,
	//			UserAvatar: v.UserAvatar,
	//		},
	//	})
	//}
	//return &types.PostInfosByUserResp{
	//	Info: infos,
	//}, nil
	return
}
