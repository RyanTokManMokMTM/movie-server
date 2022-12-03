package custom_list

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/common/pagination"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListMoviesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListMoviesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListMoviesLogic {
	return &GetListMoviesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListMoviesLogic) GetListMovies(req *types.GetListMoviesReq) (resp *types.GetListMovieResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.DAO.FindOneList(l.ctx, req.ListID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.LIST_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	limit := pagination.GetLimit(req.Limit)
	movies, err := l.svcCtx.DAO.FindListMovies(l.ctx, req.ListID, req.LastCreatedTime, int(limit))
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	listMovies := make([]types.ListMovieInfo, 0)
	for _, info := range movies {
		listMovies = append(listMovies, types.ListMovieInfo{
			CreatedTime: uint(info.CreatedAt.Unix()),
			Movies: types.MovieInfo{
				MovieID:     info.Id,
				Title:       info.Title,
				VoteAverage: info.VoteAverage,
				PosterPath:  info.PosterPath,
			},
		})
	}

	return &types.GetListMovieResp{
		ListMovies: listMovies,
	}, nil
}
