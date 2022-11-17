package movie

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieLikedCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieLikedCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieLikedCountLogic {
	return &GetMovieLikedCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieLikedCountLogic) GetMovieLikedCount(req *types.CountMovieLikesReq) (resp *types.CountMovieLikedResp, err error) {
	// todo: add your logic here and delete this line

	////get - how many likes are the movie has
	//
	//count, err := l.svcCtx.DAO.CountLikesOfMovie(l.ctx, req.MovieID)
	//if err != nil {
	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	//}
	//return &types.CountMovieLikedResp{
	//	Count: uint(count),
	//}, nil
	return
}
