package likedMovie

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLikedMovieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLikedMovieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLikedMovieLogic {
	return &DeleteLikedMovieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLikedMovieLogic) DeleteLikedMovie(req *types.DeleteLikedMoviedReq) (resp *types.DeleteLikedMovieResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	//find user
	user, err := l.svcCtx.User.FindOne(l.ctx, userID)
	if err != nil && err != sqlx.ErrNotFound {
		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("DeleteLikedMovie - user db err:%v, userID:%v", err, userID))
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if user == nil {
		//return nil, errors.Wrap(errx.NewErrCode(errx.USER_NOT_EXIST), fmt.Sprintf("DeleteLikedMovie - user db USER NOT FOUND err: %v, userID: %v", err, userID))
		return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
	}

	err = l.svcCtx.LikedMovie.Delete(l.ctx, req.MovieID)
	if err != nil {
		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("DeleteLikedMovie - user db delete err:%v, req:%+v", err, req))
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return &types.DeleteLikedMovieResp{}, nil
}
