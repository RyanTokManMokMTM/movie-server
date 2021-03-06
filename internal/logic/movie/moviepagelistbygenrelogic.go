package movie

import (
	"context"
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
	list, err := l.svcCtx.DAO.FindMovieListByGenreID(l.ctx, req.Id)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	var movies []*types.MovieInfo
	for _, v := range list.MovieInfo {
		movies = append(movies, &types.MovieInfo{
			MovieID:     v.MovieId,
			PosterPath:  v.PosterPath,
			Title:       v.Title,
			VoteAverage: v.VoteAverage,
		})
	}
	return &types.MoviePageListByGenreResponse{
		Resp: movies,
	}, nil
}

//
//func (l *MoviePageListByGenreLogic) MoviePageListByGenre(req *types.MoviePageListByGenreRequest) (resp *types.MoviePageListByGenreResponse, err error) {
//	// todo: add your logic here and delete this line
//	// return a list of movie from relation table
//	res, err := l.svcCtx.Movie.MoviePageListsByGenreID(l.ctx, req.Id, 20)
//	if err != nil && err != sqlx.ErrNotFound {
//		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("MoviePageListByGenre - movie db FIND err: %v, genreID: %v", err, req.Id))
//		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	}
//
//	var movieInfos []*types.MovieInfo
//	for _, v := range res {
//		movieInfos = append(movieInfos, &types.MovieInfo{
//			MovieId:     v.MovieId,
//			PosterPath:  v.PosterPath,
//			Title:       v.Title,
//			VoteAverage: v.VoteAverage,
//		})
//	}
//	return &types.MoviePageListByGenreResponse{
//		Resp: movieInfos,
//	}, nil
//
//}
