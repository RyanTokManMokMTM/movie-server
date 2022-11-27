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

type GetOneMovieFromUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOneMovieFromUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneMovieFromUserListLogic {
	return &GetOneMovieFromUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOneMovieFromUserListLogic) GetOneMovieFromUserList(req *types.GetOneMovieFromUserListReq) (resp *types.GetOneMovieFromUserListResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	list, err := l.svcCtx.DAO.FindOneMovieFormAnyList(l.ctx, req.MovieID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &types.GetOneMovieFromUserListResp{
				ListId:        0,
				IsMovieInList: false,
			}, nil
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if list.ListListId == 0 {
		return &types.GetOneMovieFromUserListResp{
			ListId:        0,
			IsMovieInList: false,
		}, nil
	}

	return &types.GetOneMovieFromUserListResp{
		ListId:        list.ListListId,
		IsMovieInList: true,
	}, nil

}
