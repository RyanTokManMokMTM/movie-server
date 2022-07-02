package list

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/movie-server/common/errorx"
	"github.com/ryantokmanmokmtm/movie-server/model/list"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserMovieListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserMovieListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserMovieListLogic {
	return &UpdateUserMovieListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserMovieListLogic) UpdateUserMovieList(req *types.UpdateUserListReq) (resp *types.UpdateUserListResp, err error) {
	// todo: add your logic here and delete this line
	userID := fmt.Sprintf("%v", l.ctx.Value("userID"))
	id, _ := strconv.Atoi(userID)
	_, err = l.svcCtx.User.FindOne(l.ctx, int64(id))
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	_, err = l.svcCtx.List.FindOneByUserIDAndListId(l.ctx, int64(id), req.Id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, errorx.NotFound
		}
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	updateModel := list.Lists{
		UserId:    int64(id),
		ListId:    req.Id,
		ListTitle: req.Title,
	}

	err = l.svcCtx.List.Update(l.ctx, &updateModel)
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}
	return &types.UpdateUserListResp{}, nil
}
