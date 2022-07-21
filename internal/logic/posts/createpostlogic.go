package posts

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLogic) CreatePost(req *types.CreatePostReq) (resp *types.CreatePostResp, err error) {
	// todo: add your logic here and delete this line
	//userID := ctxtool.GetUserIDFromCTX(l.ctx)
	//user, err := l.svcCtx.User.FindOne(l.ctx, userID)
	//if err != nil && err != sqlx.ErrNotFound {
	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	//}
	//
	//if user == nil {
	//	return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
	//}
	//
	//logx.Infof("UserID:%v", userID)
	//
	//postCreateTime := time.Now()
	//
	//newPost := post.Posts{
	//	PostTitle:  req.PostTitle,
	//	PostDesc:   req.PostDesc,
	//	MovieId:    req.MovieID,
	//	UserId:     userID,
	//	PostLike:   0,
	//	CreateTime: postCreateTime,
	//}
	//
	//sqlRes, err := l.svcCtx.PostModel.Insert(l.ctx, &newPost)
	//if err != nil {
	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	//}
	//
	//newPost.PostId, err = sqlRes.LastInsertId()
	//if err != nil {
	//	return nil, errx.NewErrCode(errx.DB_AFFECTED_ZERO_ERROR)
	//}
	//
	//return &types.CreatePostResp{
	//	PostID:     newPost.PostId,
	//	CreateTime: postCreateTime.Unix(),
	//}, nil
	return
}
