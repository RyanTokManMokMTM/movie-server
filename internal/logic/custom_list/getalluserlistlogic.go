package custom_list

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/common/pagination"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUserListLogic {
	return &GetAllUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllUserListLogic) GetAllUserList(req *types.AllCustomListReq) (resp *types.AllCustomListResp, err error) {
	// todo: add your logic here and delete this line

	limit := pagination.GetLimit(req.Limit)
	pageOffset := pagination.PageOffset(limit, req.Page)

	lists, count, err := l.svcCtx.DAO.FindUserLists(l.ctx, req.ID, int(limit), int(pageOffset))
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	logx.Info("total record : ", count)

	totalPage := pagination.GetTotalPageByPageSize(uint(count), limit)
	var userLists []types.ListInfo
	for _, v := range lists {
		//totalMovies, _ := l.svcCtx.DAO.CountListMovies(l.ctx, v.ListId)
		listMovies, total, _ := l.svcCtx.DAO.FindListMovies(l.ctx, v.ListId, 0, 4)

		var movieList []types.MovieInfo

		for _, movieInfo := range listMovies {
			movieList = append(movieList, types.MovieInfo{
				MovieID:     movieInfo.Id,
				Title:       movieInfo.Title,
				PosterPath:  movieInfo.PosterPath,
				VoteAverage: movieInfo.VoteAverage,
			})
		}

		userLists = append(userLists, types.ListInfo{
			ID:          v.ListId,
			Title:       v.ListTitle,
			Intro:       v.ListIntro,
			TotalMovies: uint(total),
			Movies:      movieList,
		})
	}
	return &types.AllCustomListResp{
		Lists: userLists,
		MetaData: types.MetaData{
			TotalPages:   totalPage,
			TotalResults: uint(count),
			Page:         pagination.GetPage(req.Page),
		},
	}, nil
}
