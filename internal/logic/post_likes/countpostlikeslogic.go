package post_likes

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountPostLikesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountPostLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountPostLikesLogic {
	return &CountPostLikesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountPostLikesLogic) CountPostLikes(req *types.CountPostLikesReq) (resp *types.CountPostLikesResp, err error) {
	// todo: add your logic here and delete this line
	//check post like record is exist
	//userID := ctxtool.GetUserIDFromCTX(l.ctx)
	//_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
	//	}
	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	//}

	postInfo, err := l.svcCtx.DAO.FindOnePostInfo(l.ctx, req.PostId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.POST_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//total, err := l.svcCtx.DAO.CountPostLikes(l.ctx, req.PostId)
	//if err != nil {
	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	//}
	return &types.CountPostLikesResp{
		TotalLikes: uint(len(postInfo.PostsLiked)),
	}, nil
}
