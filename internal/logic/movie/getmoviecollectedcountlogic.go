package movie

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieCollectedCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieCollectedCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieCollectedCountLogic {
	return &GetMovieCollectedCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieCollectedCountLogic) GetMovieCollectedCount(req *types.CountMovieCollectedReq) (resp *types.CountMovieCollectedResp, err error) {
	// todo: add your logic here and delete this line
	logx.Info("testing")
	count, err := l.svcCtx.DAO.CountMovieCollected(l.ctx, req.MovieID)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return &types.CountMovieCollectedResp{
		Count: uint(count),
	}, nil
}
