package likedMovie

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/movie-server/common/errorx"
	"github.com/ryantokmanmokmtm/movie-server/model/liked_movie"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLikedMovieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLikedMovieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLikedMovieLogic {
	return &CreateLikedMovieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLikedMovieLogic) CreateLikedMovie(req *types.CreateLikedMovieReq) (resp *types.CreateLikedMovieResp, err error) {
	// todo: add your logic here and delete this line
	userID := fmt.Sprintf("%v", l.ctx.Value("userID"))
	id, _ := strconv.Atoi(userID)
	_, err = l.svcCtx.User.FindOne(l.ctx, int64(id))
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	//Check movie exist ????
	//_, err = l.svcCtx.Movie.FindOne(l.ctx, req.MovieID)
	//if err != nil {
	//	if err == sqlx.ErrNotFound {
	//		return nil, errorx.NotFound
	//	}
	//	return nil, errorx.NewDefaultCodeError(err.Error())
	//}
	_, err = l.svcCtx.LikedMovie.FindOneByUserIdMovieId(l.ctx, int64(id), req.MovieID)
	if err == sqlx.ErrNotFound {
		newModel := liked_movie.LikedMovies{
			UserId:  int64(id),
			MovieId: req.MovieID,
		}

		sqlRes, err := l.svcCtx.LikedMovie.Insert(l.ctx, &newModel)
		if err != nil {
			return nil, errorx.NewDefaultCodeError(err.Error())
		}

		newModel.LikedMovieId, err = sqlRes.LastInsertId()
		if err != nil {
			return nil, errorx.NewDefaultCodeError(err.Error())
		}
		return &types.CreateLikedMovieResp{
			LikedMovieID: newModel.MovieId,
			UserID:       newModel.LikedMovieId,
		}, nil
	}

	return nil, errorx.NewDefaultCodeError("Movie has been already liked")
}
