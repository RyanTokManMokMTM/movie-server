package movie

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"gorm.io/gorm"

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
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	if userId == 0 {
		return nil, fmt.Errorf("user_id is missing")
	}

	//find that user
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not exist")
		}
		return nil, err
	}
	movie, err := l.svcCtx.DAO.FindOneMovieDetailWithUserData(l.ctx, req.MovieID, userId)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	//do we need to handle err??
	//Total Likes:
	collectedCount, _ := l.svcCtx.DAO.CountMovieCollected(l.ctx, req.MovieID)
	//Total Collected:
	likedCount, _ := l.svcCtx.DAO.CountLikesMovie(l.ctx, req.MovieID)

	var genres []types.GenreInfo
	for _, v := range movie.GenreInfo {
		genres = append(genres, types.GenreInfo{
			Id:   v.GenreId,
			Name: v.Name,
		})
	}
	//TODO: we need to include user liked and collection info
	return &types.MovieDetailResp{
		Info: types.MovieDetailInfo{
			Adult:            movie.Adult,
			BackdropPath:     movie.BackdropPath,
			MovieId:          movie.Id,
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

			//like count,collection count
			//isLike and isCollected
			IsLiked:        len(movie.Likes) == 1,
			IsCollected:    len(movie.Lists) == 1,
			CollectedCount: uint(collectedCount),
			LikedCount:     uint(likedCount),
		},
	}, nil
}
