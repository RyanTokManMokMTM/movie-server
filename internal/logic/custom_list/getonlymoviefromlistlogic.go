package custom_list

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

type GetOnlyMovieFromListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOnlyMovieFromListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnlyMovieFromListLogic {
	return &GetOnlyMovieFromListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOnlyMovieFromListLogic) GetOnlyMovieFromList(req *types.GetOneMovieFromListReq) (resp *types.GetOneMovieFromListResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	movie, err := l.svcCtx.DAO.FindOneMovieFromList(l.ctx, req.MovieID, req.ListID, userID)
	if err != nil {
		logx.Info(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.LIST_MOVIE_NOT_IN_LIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if movie.Title == "" {
		return nil, errx.NewErrCode(errx.LIST_MOVIE_NOT_IN_LIST)
	}
	return &types.GetOneMovieFromListResp{
		Movie: types.MovieInfo{
			MovieID:     movie.MovieId,
			Title:       movie.Title,
			PosterPath:  movie.PosterPath,
			VoteAverage: movie.VoteAverage,
		},
	}, nil
}
