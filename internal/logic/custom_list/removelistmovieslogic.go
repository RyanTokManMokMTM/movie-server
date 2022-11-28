package custom_list

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveListMoviesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveListMoviesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveListMoviesLogic {
	return &RemoveListMoviesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveListMoviesLogic) RemoveListMovies(req *types.RemoveListMoviesReq) (resp *types.RemoveListMoviesResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	_, err = l.svcCtx.DAO.FindOneList(l.ctx, req.ListId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.LIST_NOT_EXIST)
		}
	}

	//remove all the list
	logx.Info(req.MovieIds)
	if err := l.svcCtx.DAO.RemoveMoviesFromList(l.ctx, req.MovieIds, req.ListId, userID); err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return &types.RemoveListMoviesResp{}, nil
}
