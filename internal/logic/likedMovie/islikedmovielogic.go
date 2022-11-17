package likedMovie

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

type IsLikedMovieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsLikedMovieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsLikedMovieLogic {
	return &IsLikedMovieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsLikedMovieLogic) IsLikedMovie(req *types.IsLikedMovieReq) (resp *types.IsLikedMovieResp, err error) {
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

	//find liked movie record
	_, err = l.svcCtx.DAO.FindOneUserLikedMovie(l.ctx, req.MovieID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &types.IsLikedMovieResp{
				IsLiked: false,
			}, nil
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return &types.IsLikedMovieResp{
		IsLiked: true,
	}, nil
}
