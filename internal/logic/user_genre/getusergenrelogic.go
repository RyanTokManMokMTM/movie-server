package user_genre

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserGenreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserGenreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserGenreLogic {
	return &GetUserGenreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserGenreLogic) GetUserGenre(req *types.GetUserGenreReq) (resp *types.GetUserGenreResp, err error) {
	// todo: add your logic here and delete this line
	genres, err := l.svcCtx.DAO.FindUserGenres(l.ctx, req.UserId)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if len(*genres) == 0 {
		return &types.GetUserGenreResp{
			UserGenres: make([]types.GenreInfo, 0),
		}, nil
	}

	var genresInfo []types.GenreInfo
	for _, genre := range *genres {
		genresInfo = append(genresInfo, types.GenreInfo{
			Id:   genre.GenreId,
			Name: genre.Name,
		})
	}
	return &types.GetUserGenreResp{
		UserGenres: genresInfo,
	}, nil
}
