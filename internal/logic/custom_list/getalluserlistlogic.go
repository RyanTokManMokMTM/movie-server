package custom_list

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
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

	lists, err := l.svcCtx.DAO.FindUserLists(l.ctx, req.ID)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	var userLists []types.ListInfo
	for _, v := range lists {
		var movieList []types.MovieInfo

		for _, movieInfo := range v.MovieInfos {
			movieList = append(movieList, types.MovieInfo{
				MovieID:     movieInfo.Id,
				Title:       movieInfo.Title,
				PosterPath:  movieInfo.PosterPath,
				VoteAverage: movieInfo.VoteAverage,
			})
		}

		userLists = append(userLists, types.ListInfo{
			ID:     v.ListId,
			Title:  v.ListTitle,
			Intro:  v.ListIntro,
			Movies: movieList,
		})
	}
	return &types.AllCustomListResp{
		Lists: userLists,
	}, nil
}
