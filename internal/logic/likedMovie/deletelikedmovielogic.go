package likedMovie

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

type DeleteLikedMovieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLikedMovieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLikedMovieLogic {
	return &DeleteLikedMovieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLikedMovieLogic) DeleteLikedMovie(req *types.DeleteLikedMoviedReq) (resp *types.DeleteLikedMovieResp, err error) {
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

	if err := l.svcCtx.DAO.RemoveLikedMovie(l.ctx, req.MovieID, userID); err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return
}
