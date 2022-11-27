package user_genre

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserGenreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserGenreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserGenreLogic {
	return &UpdateUserGenreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserGenreLogic) UpdateUserGenre(req *types.UpdateUserGenreReq) (resp *types.UpdateUserGenreResp, err error) {
	// todo: add your logic here and delete this line

	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if err := l.svcCtx.DAO.UpdateUserGenres(l.ctx, req.GenreIds, userID); err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	return &types.UpdateUserGenreResp{}, nil
}
