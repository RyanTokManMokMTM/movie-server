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

func (l *MoviePageListByGenreLogic) MoviePageListByGenre(req *types.MoviePageListByGenreReq) (resp *types.MoviePageListByGenreResp, err error) {
	// todo: add your logic here and delete this line
	// return a list of movie from relation table
	list, err := l.svcCtx.DAO.FindMovieListByGenreID(l.ctx, req.Id)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	var movies []*types.MovieInfo
	for _, v := range list.MovieInfo {
		movies = append(movies, &types.MovieInfo{
			MovieID:     v.Id,
			PosterPath:  v.PosterPath,
			Title:       v.Title,
			VoteAverage: v.VoteAverage,
		})
	}
	return &types.MoviePageListByGenreResp{
		Resp: movies,
	}, nil
}
