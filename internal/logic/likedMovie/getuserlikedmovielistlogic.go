package likedMovie

import (
	"context"
	"encoding/json"
	"github.com/ryantokmanmokmtm/movie-server/common/errorx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLikedMovieListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLikedMovieListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLikedMovieListLogic {
	return &GetUserLikedMovieListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLikedMovieListLogic) GetUserLikedMovieList(req *types.AllUserLikedMoviesReq) (resp *types.AllUserAllLikedMoviesResp, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.LikedMovie.FindAllByUserIDWithMovieInfo(l.ctx, req.ID)

	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	var movieInfo []*types.LikedMovieInfo
	for _, v := range res {
		var genres []types.GenreInfo
		json.Unmarshal(v.Genres, &genres)
		movieInfo = append(movieInfo, &types.LikedMovieInfo{
			MovieID:      v.MovieId,
			MovieName:    v.MovieTitle,
			Genres:       genres,
			MoviePoster:  v.MoviePoster,
			MovieVoteAvg: v.MovieVoteAvg,
		})
	}

	return &types.AllUserAllLikedMoviesResp{
		LikedMoviesList: movieInfo,
	}, nil
}
