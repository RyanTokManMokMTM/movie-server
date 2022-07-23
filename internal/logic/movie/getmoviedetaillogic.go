package movie

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieDetailLogic {
	return &GetMovieDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieDetailLogic) GetMovieDetail(req *types.MovieDetailReq) (resp *types.MovieDetailResp, err error) {
	// todo: add your logic here and delete this line
	logx.Info("Get Movie Detail")
	movie, err := l.svcCtx.DAO.FindOneMovieDetail(l.ctx, req.MovieID)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	var genres []types.GenreInfo
	for _, v := range movie.GenreInfo {
		genres = append(genres, types.GenreInfo{
			Id:   v.GenreId,
			Name: v.Name,
		})
	}

	return &types.MovieDetailResp{
		MovieDetailInfo: types.MovieDetailInfo{
			Adult:            movie.Adult,
			BackdropPath:     movie.BackdropPath,
			MovieId:          movie.MovieId,
			OriginalLanguage: movie.OriginalLanguage,
			OriginalTitle:    movie.OriginalTitle,
			Overview:         movie.Overview,
			Popularity:       movie.Popularity,
			PosterPath:       movie.PosterPath,
			ReleaseDate:      movie.ReleaseDate,
			Title:            movie.Title,
			RunTime:          movie.RunTime,
			Video:            movie.Video,
			VoteAverage:      movie.VoteAverage,
			VoteCount:        movie.VoteCount,
			Genres:           genres,
		},
	}, nil
}
