package movie

import (
	"context"
	"errors"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MovieGenreByMovieIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMovieGenreByMovieIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MovieGenreByMovieIDLogic {
	return &MovieGenreByMovieIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MovieGenreByMovieIDLogic) MovieGenreByMovieID(req *types.MovieGenresInfoRequest) (resp *types.MovieGenresInfoResponse, err error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.Genre.FindMovieGenresByMovieID(l.ctx, req.Id)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var res []*types.GenreInfo
	for _, v := range list {
		res = append(res, &types.GenreInfo{
			Id:   v.GenreId,
			Name: v.Name,
		})
	}
	return &types.MovieGenresInfoResponse{
		Resp: res,
	}, nil
}
