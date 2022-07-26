package movie

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
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

func (l *MovieGenreByMovieIDLogic) MovieGenreByMovieID(req *types.MovieGenresInfoReq) (resp *types.MovieGenresInfoResp, err error) {
	// todo: add your logic here and delete this line
	movie, err := l.svcCtx.DAO.FindOneMovieDetail(l.ctx, req.Id)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	var genreInfo []*types.GenreInfo
	for _, v := range movie.GenreInfo {
		genreInfo = append(genreInfo, &types.GenreInfo{
			Id:   v.GenreId,
			Name: v.Name,
		})
	}

	return &types.MovieGenresInfoResp{
		Resp: genreInfo,
	}, nil
}

//
//func (l *MovieGenreByMovieIDLogic) MovieGenreByMovieID(req *types.MovieGenresInfoRequest) (resp *types.MovieGenresInfoResponse, err error) {
//	// todo: add your logic here and delete this line
//	list, err := l.svcCtx.Genre.FindMovieGenresByMovieID(l.ctx, req.Id)
//	if err != nil && err != sqlx.ErrNotFound {
//		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("MovieGenreByMovieID - genre db FIND err: %v, movieID: %v", err, req.Id))
//		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	}
//
//	if list == nil {
//		return nil, errx.NewErrCode(errx.LIST_NOT_EXIST)
//	}
//
//	var res []*types.GenreInfo
//	for _, v := range list {
//		res = append(res, &types.GenreInfo{
//			Id:   v.GenreId,
//			Name: v.Name,
//		})
//	}
//	return &types.MovieGenresInfoResponse{
//		Resp: res,
//	},
//	return
//}
