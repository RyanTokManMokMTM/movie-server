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

type LikedMovieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikedMovieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikedMovieLogic {
	return &LikedMovieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikedMovieLogic) LikedMovie(req *types.LikedMovieReq) (resp *types.LikedMovieResp, err error) {
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
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if err := l.svcCtx.DAO.CreateLikedMovie(l.ctx, req.MovieID, userID); err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return &types.LikedMovieResp{}, nil
}
