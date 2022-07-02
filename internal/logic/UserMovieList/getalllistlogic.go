package UserMovieList

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errorx"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllListLogic {
	return &GetAllListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllListLogic) GetAllList(req *types.ListsReq) (resp *types.ListsResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.List.FindAllByUpdateTimeDESC(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	var listInfos []*types.UserListInfo
	for _, v := range res {
		owner := types.UserInfo{
			Id:     v.UserId,
			Name:   v.Name,
			Email:  v.Email,
			Avatar: v.Avatar,
		}

		listInfos = append(listInfos, &types.UserListInfo{
			Id:         v.ListId,
			ListTitle:  v.ListTitle,
			Owner:      owner,
			UpdateTime: v.UpdateTime.String(),
		})
	}
	return &types.ListsResp{
		Lists: listInfos,
	}, nil
}
