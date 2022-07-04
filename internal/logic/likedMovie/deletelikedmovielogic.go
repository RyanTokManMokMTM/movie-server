package likedMovie

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/movie-server/common/errorx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"

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
	userID := fmt.Sprintf("%v", l.ctx.Value("userID"))
	id, _ := strconv.Atoi(userID)
	_, err = l.svcCtx.User.FindOne(l.ctx, int64(id))
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	res, err := l.svcCtx.LikedMovie.FindOneByUserIdMovieId(l.ctx, int64(id), req.MovieID)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, errorx.NotFound
		}
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	err = l.svcCtx.LikedMovie.Delete(l.ctx, res.LikedMovieId)
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	return &types.DeleteLikedMovieResp{}, nil
}
