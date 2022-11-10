package likedMovie

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
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	limit := pagination.GetLimit(req.Limit)
	pageOffset := pagination.PageOffset(pagination.DEFAULT_PAGE_SIZE, req.Page)

	u, count, err := l.svcCtx.DAO.GetUserLikedMovies(l.ctx, req.ID, int(limit), int(pageOffset))
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	logx.Info("total record : ", count)

	totalPage := pagination.GetTotalPageByPageSize(uint(count), pagination.DEFAULT_PAGE_SIZE)

	var likedMovie []*types.LikedMovieInfo
	for _, v := range u.MovieInfos {
		var genres []types.GenreInfo
		for _, t := range v.GenreInfo {
			genres = append(genres, types.GenreInfo{
				Id:   t.GenreId,
				Name: t.Name,
			})
		}

		likedMovie = append(likedMovie, &types.LikedMovieInfo{
			MovieID:      v.Id,
			MovieName:    v.Title,
			Genres:       genres,
			MoviePoster:  v.PosterPath,
			MovieVoteAvg: v.VoteAverage,
		})
	}
	return &types.AllUserAllLikedMoviesResp{
		LikedMoviesList: likedMovie,
		MetaData: types.MetaData{
			TotalPages:   totalPage,
			TotalResults: uint(count),
			Page:         pagination.GetPage(req.Page),
		},
	}, nil
}
