package likedMovie

import (
	"context"
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
	//userID := ctxtool.GetUserIDFromCTX(l.ctx)
	//
	////find user
	//user, err := l.svcCtx.User.FindOne(l.ctx, userID)
	//if err != nil && err != sqlx.ErrNotFound {
	//	return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("CreateLikedMovie - user db err:%v, userID:%v", err, userID))
	//}
	//
	//if user == nil {
	//	return nil, errors.Wrap(errx.NewErrCode(errx.USER_NOT_EXIST), fmt.Sprintf("CreateLikedMovie - user db USER NOT FOUND err: %v, userID: %v", err, userID))
	//}
	//
	////Check movie exist ????
	////_, err = l.svcCtx.Movie.FindOne(l.ctx, req.MovieId)
	////if err != nil {
	////	if err == sqlx.ErrNotFound {
	////		return nil, errorx.NotFound
	////	}
	////	return nil, errorx.NewDefaultCodeError(err.Error())
	////}
	//likedMovie, err := l.svcCtx.LikedMovie.FindOneByUserIdMovieId(l.ctx, userID, req.MovieID)
	//if likedMovie == nil {
	//	newModel := liked_movie.LikedMovies{
	//		UserId:  userID,
	//		MovieId: req.MovieID,
	//	}
	//
	//	sqlRes, err := l.svcCtx.LikedMovie.Insert(l.ctx, &newModel)
	//	if err != nil {
	//		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("CreateLikedMovie - LikedMovie db Insert err: %v, req: %+v", err, req))
	//		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	//	}
	//
	//	newModel.LikedMovieId, err = sqlRes.LastInsertId()
	//	if err != nil {
	//		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_AFFECTED_ZERO_ERROR), fmt.Sprintf("CreateLikedMovie - LikedMovie db Insert.LastInsertId err: %v, req: %+v", err, req))
	//		return nil, errx.NewErrCode(errx.DB_AFFECTED_ZERO_ERROR)
	//	}
	//	return &types.CreateLikedMovieResp{
	//		LikedMovieID: newModel.MovieId,
	//		UserID:       newModel.LikedMovieId,
	//	}, nil
	//}
	//
	////return nil, errors.Wrap(errx.NewErrCode(errx.MOVIE_ALREADY_LIKED), fmt.Sprintf("CreateLikedMovie - LikedMovie db FIND err: %v, req: %+v", err, req))
	//return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	return
}
