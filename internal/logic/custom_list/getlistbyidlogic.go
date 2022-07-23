package custom_list

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetListByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListByIDLogic {
	return &GetListByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListByIDLogic) GetListByID(req *types.UserListReq) (resp *types.UserListResp, err error) {
	// todo: add your logic here and delete this line

	list, err := l.svcCtx.DAO.FindOneList(l.ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.LIST_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	var movies []types.MovieInfo
	for _, v := range list.MovieInfos {
		movies = append(movies, types.MovieInfo{
			MovieID:     v.MovieId,
			Title:       v.Title,
			PosterPath:  v.PosterPath,
			VoteAverage: v.VoteAverage,
		})
	}

	return &types.UserListResp{
		List: types.ListInfo{
			ID:     list.ListId,
			Title:  list.ListTitle,
			Movies: movies,
			//List of movie????
		},
	}, nil
}

//
//func (l *GetListByIDLogic) GetListByID(req *types.UserListReq) (resp *types.UserListResp, err error) {
//	// todo: add your logic here and delete this line
//	//res, err := l.svcCtx.List.FindOne(l.ctx, req.ID)
//	//if err != nil && err != sqlx.ErrNotFound {
//	//	//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("GetListByID - List db err: %v, ListID: %v", err, req.ID))
//	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	//}
//	//
//	//if res == nil {
//	//	//return nil, errors.Wrap(errx.NewErrCode(errx.LIST_NOT_EXIST), fmt.Sprintf("GetListByID - List db FIND NOT FOUND err: %v, ListID: %v", err, req.ID))
//	//	return nil, errx.NewErrCode(errx.LIST_NOT_EXIST)
//	//}
//	//
//	//return &types.UserListResp{
//	//	List: types.ListInfo{
//	//		ID:       res.ListId,
//	//		Title:    res.ListTitle,
//	//		UpdateOn: res.UpdateTime.Unix(),
//	//	},
//	//}, nil
//	return
//}
