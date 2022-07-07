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

type UpdateCustomListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCustomListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCustomListLogic {
	return &UpdateCustomListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCustomListLogic) UpdateCustomList(req *types.UpdateCustomListReq) (resp *types.UpdateCustomListResp, err error) {
	// todo: add your logic here and delete this line
	userID := fmt.Sprintf("%v", l.ctx.Value("userID"))
	id, _ := strconv.Atoi(userID)
	_, err = l.svcCtx.User.FindOne(l.ctx, int64(id))
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	res, err := l.svcCtx.List.FindOne(l.ctx, req.ID)
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	//title is a required field
	res.ListTitle = req.Title

	err = l.svcCtx.List.Update(l.ctx, res)
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}
	return &types.UpdateCustomListResp{}, nil
}
