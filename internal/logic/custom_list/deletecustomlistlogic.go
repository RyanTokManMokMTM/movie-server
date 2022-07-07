package custom_list

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/movie-server/common/errorx"
	"strconv"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCustomListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCustomListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCustomListLogic {
	return &DeleteCustomListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCustomListLogic) DeleteCustomList(req *types.DeleteCustomListReq) (resp *types.DeleteCustomListResp, err error) {
	// todo: add your logic here and delete this line
	userID := fmt.Sprintf("%v", l.ctx.Value("userID"))
	id, _ := strconv.Atoi(userID)
	_, err = l.svcCtx.User.FindOne(l.ctx, int64(id))
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	_, err = l.svcCtx.List.FindOneByUserIDAndListId(l.ctx, int64(id), req.ID)
	if err != nil {
		return nil, errorx.NewDefaultCodeError("You are not allow to delete the list")
	}

	err = l.svcCtx.List.Delete(l.ctx, req.ID)
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	return &types.DeleteCustomListResp{}, nil
}
