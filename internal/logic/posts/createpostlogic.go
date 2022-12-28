package posts

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"gorm.io/gorm"

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
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	//find that user
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	newPost := models.Post{
		PostTitle:   req.PostTitle,
		PostDesc:    req.PostDesc,
		MovieInfoId: req.MovieID,
		UserId:      userID,
	}

	if err := l.svcCtx.DAO.CreateNewPost(l.ctx, &newPost); err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return &types.CreatePostResp{
		PostID:     newPost.PostId,
		CreateTime: newPost.CreatedAt.Unix(),
	}, nil
}

//
//func (l *CreatePostLogic) CreatePost(req *types.CreatePostReq) (resp *types.CreatePostResp, err error) {
//	// todo: add your logic here and delete this line
//	//userID := ctxtool.GetUserIDFromCTX(l.ctx)
//	//user, err := l.svcCtx.User.FindOne(l.ctx, userID)
//	//if err != nil && err != sqlx.ErrNotFound {
//	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	//}
//	//
//	//if user == nil {
//	//	return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
//	//}
//	//
//	//logx.Infof("UserID:%v", userID)
//	//
//	//postCreateTime := time.Now()
//	//
//	//newPost := post.Posts{
//	//	PostTitle:  req.PostTitle,
//	//	PostDesc:   req.PostDesc,
//	//	ID:    req.MovieID,
//	//	UserId:     userID,
//	//	PostLike:   0,
//	//	CreateTime: postCreateTime,
//	//}
//	//
//	//sqlRes, err := l.svcCtx.PostModel.CreateOne(l.ctx, &newPost)
//	//if err != nil {
//	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	//}
//	//
//	//newPost.PostId, err = sqlRes.LastInsertId()
//	//if err != nil {
//	//	return nil, errx.NewErrCode(errx.DB_AFFECTED_ZERO_ERROR)
//	//}
//	//
//	//return &types.CreatePostResp{
//	//	PostID:     newPost.PostId,
//	//	CreateTime: postCreateTime.Unix(),
//	//}, nil
//	return
//}
