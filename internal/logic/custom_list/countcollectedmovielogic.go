package custom_list

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountCollectedMovieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountCollectedMovieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountCollectedMovieLogic {
	return &CountCollectedMovieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountCollectedMovieLogic) CountCollectedMovie(req *types.CountCollectedMovieReq) (resp *types.CountCollectedMovieResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	count, err := l.svcCtx.DAO.CountCollectedMovie(l.ctx, req.UserID)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return &types.CountCollectedMovieResp{
		Total: uint(count),
	}, nil
}
