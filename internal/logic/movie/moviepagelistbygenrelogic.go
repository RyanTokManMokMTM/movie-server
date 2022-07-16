package movie

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MoviePageListByGenreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoviePageListByGenreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoviePageListByGenreLogic {
	return &MoviePageListByGenreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoviePageListByGenreLogic) MoviePageListByGenre(req *types.MoviePageListByGenreRequest) (resp *types.MoviePageListByGenreResponse, err error) {
	// todo: add your logic here and delete this line
	// return a list of movie from relation table
	res, err := l.svcCtx.Movie.MoviePageListsByGenreID(l.ctx, req.Id, 20)
	if err != nil {
		return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("MoviePageListByGenre - movie db FIND err: %v, genreID: %v", err, req.Id))
	}

	var movieInfos []*types.MovieInfo
	for _, v := range res {
		movieInfos = append(movieInfos, &types.MovieInfo{
			MovieID:     v.MovieId,
			PosterPath:  v.PosterPath,
			Title:       v.Title,
			VoteAverage: v.VoteAverage,
		})
	}
	return &types.MoviePageListByGenreResponse{
		Resp: movieInfos,
	}, nil

}
