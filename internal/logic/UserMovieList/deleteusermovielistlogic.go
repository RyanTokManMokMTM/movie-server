package UserMovieList

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/movie-server/common/errorx"
	"strconv"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserMovieListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserMovieListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserMovieListLogic {
	return &DeleteUserMovieListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserMovieListLogic) DeleteUserMovieList(req *types.DeleteUserListReq) (resp *types.DeleteListDetailInfoResp, err error) {
	// todo: add your logic here and delete this line
	userID := fmt.Sprintf("%v", l.ctx.Value("userID"))
	id, _ := strconv.Atoi(userID)
	_, err = l.svcCtx.User.FindOne(l.ctx, int64(id))
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	err = l.svcCtx.List.Delete(l.ctx, req.Id)
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	return &types.DeleteListDetailInfoResp{}, nil
}
