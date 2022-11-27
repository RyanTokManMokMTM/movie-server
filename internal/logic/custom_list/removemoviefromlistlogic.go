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

type RemoveMovieFromListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveMovieFromListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveMovieFromListLogic {
	return &RemoveMovieFromListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveMovieFromListLogic) RemoveMovieFromList(req *types.RemoveMovieReq) (resp *types.RemoveMovieResp, err error) {
	// todo: add your logic here and delete this line
	logx.Infof("REMOVE MOVIE FROM LIST : listID: %v, movieID: %v", req.ListID, req.MovieID)
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if err := l.svcCtx.DAO.RemoveMovieFromList(l.ctx, req.MovieID, req.ListID, userID); err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	return &types.RemoveMovieResp{}, nil
}
