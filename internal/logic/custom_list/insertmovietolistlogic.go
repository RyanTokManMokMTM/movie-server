package custom_list

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertMovieToListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertMovieToListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertMovieToListLogic {
	return &InsertMovieToListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertMovieToListLogic) InsertMovieToList(req *types.InsertMovieReq) (resp *types.InsertMovieResp, err error) {
	// todo: add your logic here and delete this line

	logx.Infof("INSERT MOVIE TO LIST: LIST ID :%v, MOVIE:ID:%v", req.ListID, req.MovieID)

	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//Check movie exists
	_, err = l.svcCtx.DAO.FindOneMovie(l.ctx, req.MovieID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.MOVIE_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//Movie is already add to any List?

	_, err = l.svcCtx.DAO.FindOneMovieFormAnyList(l.ctx, req.MovieID, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := l.svcCtx.DAO.InsertMovieToList(l.ctx, req.MovieID, req.ListID, userID); err != nil {
			return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
		}

		return &types.InsertMovieResp{}, nil
	}

	return nil, errx.NewErrCode(errx.LIST_MOVIE_ALREADY_IN_LIST)

}
